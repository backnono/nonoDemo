package observability

import (
	"context"
)

type eventKey struct{}

type Event struct {
	values map[string]interface{}
}

// AppendEvent 向上下文中追加事件信息
func AppendEvent(ctx context.Context, key string, e interface{}) context.Context {
	event := ctx.Value(eventKey{})
	if event == nil {
		event = &Event{values: map[string]interface{}{}}
		ctx = context.WithValue(ctx, eventKey{}, event)
		event.(*Event).values[key] = e
		return ctx
	}
	if _event, ok := event.(*Event); ok {
		_event.values[key] = e
		return ctx
	}
	event = &Event{values: map[string]interface{}{}}
	ctx = context.WithValue(ctx, eventKey{}, event)
	event.(*Event).values[key] = e
	return ctx
}

func AppendEvents(ctx context.Context, args ...interface{}) context.Context {
	event := ctx.Value(eventKey{})
	if event == nil {
		event = &Event{values: map[string]interface{}{}}
		ctx = context.WithValue(ctx, eventKey{}, event)
		for i := 0; i < len(args); {
			event.(*Event).values[args[i].(string)] = args[i+1]
			i += 2
		}
		return ctx
	}
	if _event, ok := event.(*Event); ok {
		for i := 0; i < len(args); {
			_event.values[args[i].(string)] = args[i+1]
			i += 2
		}
		return ctx
	}
	event = &Event{values: map[string]interface{}{}}
	ctx = context.WithValue(ctx, eventKey{}, event)
	for i := 0; i < len(args); {
		event.(*Event).values[args[i].(string)] = args[i+1]
		i += 2
	}
	return ctx
}

// GetEvent 从上下文中获取事件信息
func GetEvent(ctx context.Context) *Event {
	event := ctx.Value(eventKey{})
	if event == nil {
		return nil
	}
	if _event, ok := event.(*Event); ok {
		return _event
	}
	return nil
}

func GetValue(ctx context.Context, key string) interface{} {
	event := GetEvent(ctx)
	if event == nil {
		return nil
	}
	return event.GetValue(key)
}

func (e *Event) EventToArgList() []interface{} {
	var args []interface{}
	for k, v := range e.values {
		args = append(args, k, v)
	}
	return args
}

func (e *Event) GetValue(key string) interface{} {
	if v, exist := e.values[key]; exist {
		return v
	} else {
		return nil
	}
}

func (e *Event) Append(args ...interface{}) {
	for i := 0; i < len(args); {
		e.values[args[i].(string)] = args[i+1]
		i += 2
	}
}
