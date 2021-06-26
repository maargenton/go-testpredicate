package impl

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/maargenton/go-testpredicate/pkg/internal/predicate"
)

// All tests if all values of a collection match the given predicate
func All(p *predicate.Predicate) (desc string, f predicate.PredicateFunc) {
	desc = fmt.Sprintf("∀ x ∈ value, %v", p.FormatDescription("x"))
	f = func(v interface{}) (r bool, ctx []predicate.ContextValue, err error) {
		vv := reflect.ValueOf(v)
		switch vv.Kind() {
		case reflect.Array, reflect.Slice:
			for i := 0; i < vv.Len(); i++ {
				success, ctx := p.Evaluate(vv.Index(i).Interface())
				if !success {
					return false, updateContext(ctx, i), nil
				}
			}

		default:
			return false, nil, fmt.Errorf(
				"value of type '%v' is not a collection",
				vv.Type())
		}
		return true, nil, nil
	}
	return
}

// Any tests if at least one values of a collection match the given predicate
func Any(p *predicate.Predicate) (desc string, f predicate.PredicateFunc) {
	desc = fmt.Sprintf("∃ x ∈ value, %v", p.FormatDescription("x"))
	f = func(v interface{}) (r bool, ctx []predicate.ContextValue, err error) {
		vv := reflect.ValueOf(v)
		switch vv.Kind() {
		case reflect.Array, reflect.Slice:
			for i := 0; i < vv.Len(); i++ {
				success, subctx := p.Evaluate(vv.Index(i).Interface())
				if success {
					return true, nil, nil
				}
				if len(ctx) == 0 {
					ctx = updateContext(subctx, i)
				}
			}

		default:
			return false, nil, fmt.Errorf(
				"value of type '%v' is not a collection",
				vv.Type())
		}
		return false, ctx, nil
	}
	return
}

func updateContext(
	ctx []predicate.ContextValue, index int) (
	result []predicate.ContextValue) {

	for _, v := range ctx {
		if v.Name == "expected" {
			continue
		}
		v.Name = formatName(v.Name, index)
		result = append(result, v)
	}
	return result
}

func formatName(name string, index int) string {
	if i := strings.Index(name, "@("); i > 0 {
		j := strings.Index(name[i+2:], ")")
		basename := name[:i]
		indexes := name[i+2 : i+2+j]
		indexList := strings.Split(indexes, ",")
		indexList = append([]string{fmt.Sprintf("%v", index)}, indexList...)
		return basename + fmt.Sprintf("@(%v)", strings.Join(indexList, ","))

	}
	return name + fmt.Sprintf(" @(%v)", index)

}
