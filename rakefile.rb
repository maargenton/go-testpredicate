# =============================================================================
#
# MODULE      : rakefile.rb
# PROJECT     : go-testpredicate
# DESCRIPTION :
#
# Copyright (c) 2016-2021, Marc-Antoine Argenton.  All rights reserved.
# =============================================================================

require 'fileutils'

task default: [:build]


desc 'Display build information'
task :info do
    summary = {
        "Module" =>  GoBuild.default.gomod,
        "Version" => GoBuild.default.version,
        "Source" =>  File.join(BuildInfo.default.remote, "tree", BuildInfo.default.commit[0,10]),
    }

    if GoBuild.default.targets.count > 0 then
        summary["Main target"] = GoBuild.default.main_target

        if GoBuild.default.targets.count > 1 then
            targets = (GoBuild.default.targets.keys - [GoBuild.default.main_target])
            summary["Additional targets"] = targets
        end
    end

    # record_summary("## Build summary\n\n#{format_summary_table(summary)}\n")

    print_summary(summary)
    export_summary("## Build summary\n\n#{format_summary_table(summary)}\n")
    export_env("VERSION=#{GoBuild.default.version}")
end


desc 'Display inferred build version string'
task :version do
    puts GoBuild.default.version
end


desc 'Run all tests and capture results'
task :test => [:info] do
    FileUtils.makedirs( ['./build/artifacts'] )
    success = go_test()
    go_testreport('build/go-test-result.json',
        '--md-shift-headers=1',
        '-oyaml=build/artifacts/test-report.yaml',
        '-omd=build/artifacts/test-report.md',
        '-omdsf=build/go-test-summary.md',
        '-omdsfd=build/go-test-details.md',
    )

    puts File.read('build/go-test-summary.md')
    export_summary(File.read('build/go-test-details.md'))

    exit(1) if !success
end

desc 'Build and publish both release archive and associated container image'
task :build => [:info, :test] do
    # Nothing to do here
    generate_release_notes()
end

desc 'Remove build artifacts'
task :clean do
    FileUtils.rm_rf('./build')
end

def generate_release_notes()
    version = BuildInfo.default.version
    File.write( 'build/release_notes.md',  extract_release_notes(version,
        # prefix: "go-testreport",
        input:'RELEASES.md',
    ))
end



# ----------------------------------------------------------------------------
# BuildInfo : Helper to extract version inforrmation for git repo
# ----------------------------------------------------------------------------

class BuildInfo
    class << self
        def default() return @default ||= new end
    end

    def initialize()
        if _git('rev-parse --is-shallow-repository') == 'true'
            puts "Fetching missing information from remote ..."
            system(' git fetch --prune --tags --unshallow')
        end
    end

    def name()      return @name    ||= _name()     end
    def version()   return @version ||= _version()  end
    def remote()    return @remote  ||= _remote()   end
    def commit()    return @commit  ||= _commit()   end
    def dir()       return @dir     ||= _dir()      end

    private
    def _git( cmd ) return `git #{cmd} 2>/dev/null`.strip()     end
    def _commit()   return _git('rev-parse HEAD')               end
    def _dir()      return _git('rev-parse --show-toplevel')    end

    def _name()
        remote_basename = File.basename(remote() || "" )
        return remote_basename if remote_basename != ""
        return File.basename(File.expand_path("."))
    end

    def _version()
        v, b, n, g = _info()                    # Extract base info from git branch and tags
        m = _mtag()                             # Detect locally modified files
        v = _patch(v) if n > 0 || !m.nil?       # Increment patch if needed to to preserve semver orderring
        b = 'rc' if _is_default_branch(b, v)    # Rename branch to 'rc' for default branch
        return v if b == 'rc' && n == 0 && m.nil?
        return "#{v}-" + [b, n, g, m].compact().join('.')
    end

    def _info()
        # Note: Due to glob(7) limitations, the following pattern enforces
        # 3-part dot-separated sequences starting with a digit,
        # rather than 3 dot-separated numbers.
        d = _git("describe --always --tags --long  --match 'v[0-9]*.[0-9]*.[0-9]*'").strip.split('-')
        if d.count != 0
            b = _git("rev-parse --abbrev-ref HEAD").strip.gsub(/[^A-Za-z0-9\._-]+/, '-')
            return ['v0.0.0', b, _git("rev-list --count HEAD").strip.to_i, "g#{d[0]}"] if d.count == 1
            return [d[0], b, d[1].to_i, d[2]] if d.count == 3
        end
        return ['v0.0.0', "none", 0, 'g0000000']
    end

    def _is_default_branch(b, v)
        # Check branch name against common main branch names, and branch name
        # that matches the beginning of the version strings e.g. 'v1' is
        # considered a default branch for version 'v1.x.y'.
        return ["main", "master", "HEAD"].include?(b) ||
            (!v.nil? && v.start_with?(b))
    end

    def _patch(v)
        # Increment the patch number by 1, so that intermediate version strings
        # sort between the last tag and the next tag according to semver.
        #   v0.6.1
        #       v0.6.1-maa-cleanup.1.g6ede8cd   <-- with _patch()
        #   v0.6.0
        #       v0.6.0-maa-cleanup.1.g6ede8cd   <-- without _patch()
        #   v0.5.99
        vv = v[1..-1].split('.').map { |v| v.to_i }
        vv[-1] += 1
        v = "v" + vv.join(".")
        return v
    end

    def _mtag()
        # Generate a `.mXXXXXXXX` fragment based on latest mtime of modified
        # files in the index. Returns `nil` if no files are locally modified.
        status = _git("status --porcelain=2 --untracked-files=no")
        files = status.lines.map {|l| l.strip.split(/ +/).last }.map { |n| n.split(/\t/).first }
        t = files.map { |f| File.mtime(f).to_i rescue nil }.compact.max
        return t.nil? ? nil : "m%08x" % t
    end

    GIT_SSH_REPO = /git@(?<host>[^:]+):(?<path>.+).git/
    def _remote()
        remote = _git('remote get-url origin')
        m = GIT_SSH_REPO.match(remote)
        return remote if m.nil?

        host = m[:host]
        host = "github.com" if host.end_with? ("github.com")
        return "https://#{host}/#{m[:path]}/"
    end
end



# ----------------------------------------------------------------------------
# GoBuild : Helper to build go projects
# ----------------------------------------------------------------------------

class GoBuild
    class << self
        def default() return @default ||= new end
    end

    def initialize( buildinfo = nil )
        @buildinfo = buildinfo || BuildInfo.default
    end

    def gomod()         return @gomod       ||= _gomod()            end
    def targets()       return @tagets      ||= _targets()          end
    def main_target()   return @main_target ||= _main_target()      end
    def version()       return @version     ||= @buildinfo.version  end
    def ldflags()       return @ldflags     ||= _ldflags()          end

    def commands(action = 'build')
        flags = %Q{"#{ldflags}"}
        Hash[targets.map do |name, input|
            output = File.join( './build/bin', name )
            cmd = [ "go #{action} -trimpath -ldflags #{flags}",
                ("-o #{output}" if action == 'build'),
                "#{input}"
            ].compact.join(' ')
            [name, cmd]
        end]
    end

private
    def _gomod()
        return '' if !File.readable?('go.mod')
        File.foreach('go.mod') do |l|
            return l[7..-1].strip if l.start_with?( 'module ' )
        end
    end

    def _targets()
        targets = Hash[Dir["./cmd/**/main.go"].map do |f|
            path = File.dirname(f)
            [File.basename(path), File.join( path, "..." )]
        end]
        targets[File.basename(gomod)] = "." if File.exist?("./main.go")
        targets
    end

    def _ldflags()
        prefix = "#{gomod}/pkg/buildinfo"
        {   Version: @buildinfo.version,
            GitHash: @buildinfo.commit,
            GitRepo: @buildinfo.remote,
            BuildRoot: @buildinfo.dir
        }.map { |k,v| "-X #{prefix}.#{k}=#{v}"}.join(' ')
    end

    def _main_target()
        mod = File.basename(gomod)
        targets.keys.min_by{ |v| _lev(v, mod)}
    end

    def _lev(a, b, memo={})
        return b.size if a.empty?
        return a.size if b.empty?
        return memo[[a, b]] ||= [
            _lev(a.chop, b, memo) + 1,
            _lev(a, b.chop, memo) + 1,
            _lev(a.chop, b.chop, memo) + (a[-1] == b[-1] ? 0 : 1)
        ].min
    end
end

def go_test()
    FileUtils.makedirs( ['./build'] )
    cmd = "go test -race " +
        "-coverprofile=build/go-test-coverage.txt -covermode=atomic " +
        "-json ./... > build/go-test-result.json"
    puts cmd
    system(cmd)
end

def go_testreport(*args)
    cmd = %w{go run github.com/maargenton/go-testreport@v0.1.6}
    cmd += args
    puts cmd
    system(*cmd)
end



# ----------------------------------------------------------------------------
# DockerHelper : Helper to build go projects
# ----------------------------------------------------------------------------

def docker_registry_tags(base_tag)
    return [github_registry_tag(base_tag)]
end

def github_registry_tag(base_tag)
    return if ENV['GITHUB_ACTOR'].nil? || ENV['GITHUB_REPOSITORY'].nil?
    if ENV['GITHUB_TOKEN'].nil? then
        puts "Found GitHub Actiona context but no 'GITHUB_TOKEN'."
        puts "Image will not be pushed to GitHub package registry."
        puts "To resolve this issue, add the following to your workflow:"
        puts "  env:"
        puts "    GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}"
        return
    end
    # Authenticate
    puts "Authenticating with docker.pkg.github.com..."
    system("echo ${GITHUB_TOKEN} | docker login ghcr.io --username ${GITHUB_ACTOR} --password-stdin")
    puts "Failed to authenticate with docker.pkg.github.com" if $?.exitstatus != 0

    return File.join('ghcr.io', ENV['GITHUB_REPOSITORY_OWNER'], base_tag)
end



# ----------------------------------------------------------------------------
# Github Actions integration
# ----------------------------------------------------------------------------

def export_summary(content)
    return if ENV['GITHUB_STEP_SUMMARY'].nil?
    summary_filename = ENV['GITHUB_STEP_SUMMARY']
    open(summary_filename, 'a') do |f|
        f.puts content
    end
end

def export_env(env)
    return if ENV['GITHUB_ENV'].nil?
    env_filename = ENV['GITHUB_ENV']
    open(env_filename, 'a') do |f|
        f.puts env
    end
end



# ----------------------------------------------------------------------------
# Build summary generator
# ----------------------------------------------------------------------------

# def record_summary(content)
#     return if ENV['GITHUB_STEP_SUMMARY'].nil?
#     summary_filename = ENV['GITHUB_STEP_SUMMARY']
#     open(summary_filename, 'a') do |f|
#         f.puts content
#     end
# end

def print_summary(summary)
    puts "---"
    width = summary.select { |k,v| !_single_value(v).nil? }.map { |k,v| k.length }.max
    summary.each do |k,v|
        vv = _single_value(v)
        if vv.nil?
            puts "#{k}:"
            puts v.map { |t| "  - #{t}" }.join(" \n")
        else
            puts "#{(k+':').ljust(width+1)} #{vv}"
        end
    end
    puts "---"
end

def _single_value(value)
    if value.respond_to?(:each) && value.respond_to?(:count)
        return nil if value.count > 1
        return value[0]
    end
    return value
end

def format_summary_table(summary)
    o = "| | |\n|-|-|\n"
    summary.each do |key, value|
        if value.respond_to?('each')
            value.each_with_index do |v, i|
                o += (i == 0) ? "| #{key} | `#{v}`\n" : "| | `#{v}`\n"
            end
        else
            o += "| #{key} | `#{value}`\n"
        end
    end
    return o
end



# ----------------------------------------------------------------------------
# Release notes generator
# ----------------------------------------------------------------------------

def extract_release_notes(version, prefix:nil, input:nil, checksums:nil)
    rn = ""
    rn += "#{prefix} #{version}\n\n" if prefix
    rn += load_release_notes(input, version) if input
    rn += "\n## Checksums\n\n```\n" + File.read(checksums) + "```\n" if checksums
    rn
end

def load_release_notes(filename, version)
    notes, capture = [], false
    File.readlines(filename).each do |l|
        if l.start_with?( "# ")
            break if capture
            capture = true if version.start_with?(l[2..-1].strip())
        elsif capture
            notes << l
        end
    end
    notes.shift while (notes.first || "-").strip == ""
    return notes.join()
end



# ----------------------------------------------------------------------------
# Definitions to help formating 'rake watch' results
# ----------------------------------------------------------------------------

TERM_WIDTH = `tput cols`.to_i || 80

def tty_red(str);           "\e[31m#{str}\e[0m" end
def tty_green(str);         "\e[32m#{str}\e[0m" end
def tty_blink(str);         "\e[5m#{str}\e[25m" end
def tty_reverse_color(str); "\e[7m#{str}\e[27m" end

def print_separator( success = true )
    if success
        puts tty_green( "-" * TERM_WIDTH )
    else
        puts tty_reverse_color(tty_red( "-" * TERM_WIDTH ))
    end
end



# ----------------------------------------------------------------------------
# Definition of watch task, that monitors the project folder for any relevant
# file change and runs the unit test of the project.
# ----------------------------------------------------------------------------

def watch( *glob )
    yield unless block_given?
    files = []
    loop do
        new_files = Dir[*glob].map {|file| File.mtime(file) }
        yield if new_files != files
        files = new_files
        sleep 0.5
    end
end

# task :watch do
#     watch( '**/*.{c,cc,cpp,h,hh,hpp,ld}', 'Makefile' ) do
#         success = system "clear && rake"
#         print_separator( success )
#     end
# end
