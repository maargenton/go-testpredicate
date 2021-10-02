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

desc 'Display build information, including inferred verison number that uniquely identies the build product'
task :info do
    puts "Name:    #{BuildInfo.default.name}"
    puts "Version: #{BuildInfo.default.version}"
    puts "Remote:  #{BuildInfo.default.remote}"
    puts "Commit:  #{BuildInfo.default.commit}"
    puts "Dir:     #{BuildInfo.default.dir}"
end

desc 'Display inferred build version number'
task :version do
    puts BuildInfo.default.version
end

desc 'Build and publish both release archive and associated container image'
task :build do
    FileUtils.makedirs( ['./build/artifacts'] )
    version = BuildInfo.default.version
    puts "Version: #{version}"
    File.write( 'build/release_notes',  generate_release_notes(version,
        # prefix: "go-testpredicate",
        input:'RELEASES.md',
    ))
end

desc 'Remove build artifacts'
task :clean do
    FileUtils.rm_rf('./build')
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
# Release notes generator
# ----------------------------------------------------------------------------

def generate_release_notes(version, prefix:nil, input:nil, checksums:nil)
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
