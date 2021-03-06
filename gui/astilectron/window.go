package astilectron

import (
	"context"
	"net/url"

	"github.com/asticode/go-astitools/context"
	"github.com/asticode/go-astitools/url"
	"github.com/pkg/errors"
)

// Window errors
var (
	ErrWindowDestroyed = errors.New("window.destroyed")
)

// Window represents a window
// https://github.com/electron/electron/blob/master/docs/api/browser-window.md
// TODO Add missing window options
// TODO Add missing window methods
// TODO Add missing window events
type Window struct {
	cancel context.CancelFunc
	ctx    context.Context
	c      *asticontext.Canceller
	d      *dispatcher
	id     string
	o      *WindowOptions
	url    *url.URL
	w      *writer
}

// WindowOptions represents window options
// We must use pointers since GO doesn't handle optional fields whereas NodeJS does. Use PtrBool, PtrInt or PtrStr
// to fill the struct
type WindowOptions struct {
	AcceptFirstMouse       *bool   `json:"acceptFirstMouse,omitempty"`
	AlwaysOnTop            *bool   `json:"alwaysOnTop,omitempty"`
	AutoHideMenuBar        *bool   `json:"autoHideMenuBar,omitempty"`
	BackgroundColor        *string `json:"backgroundColor,omitempty"`
	Center                 *bool   `json:"center,omitempty"`
	Closable               *bool   `json:"closable,omitempty"`
	DisableAutoHideCursor  *bool   `json:"disableAutoHideCursor,omitempty"`
	EnableLargerThanScreen *bool   `json:"enableLargerThanScreen,omitempty"`
	Focusable              *bool   `json:"focusable,omitempty"`
	Frame                  *bool   `json:"frame,omitempty"`
	Fullscreen             *bool   `json:"fullscreen,omitempty"`
	Fullscreenable         *bool   `json:"fullscreenable,omitempty"`
	HasShadow              *bool   `json:"hasShadow,omitempty"`
	Height                 *int    `json:"height,omitempty"`
	Icon                   *string `json:"icon,omitempty"`
	Kiosk                  *bool   `json:"kiosk,omitempty"`
	MaxHeight              *int    `json:"maxHeight,omitempty"`
	Maximizable            *bool   `json:"maximizable,omitempty"`
	MaxWidth               *int    `json:"maxWidth,omitempty"`
	MinHeight              *int    `json:"minHeight,omitempty"`
	Minimizable            *bool   `json:"minimizable,omitempty"`
	MinWidth               *int    `json:"minWidth,omitempty"`
	Modal                  *bool   `json:"modal,omitempty"`
	Movable                *bool   `json:"movable,omitempty"`
	Resizable              *bool   `json:"resizable,omitempty"`
	Show                   *bool   `json:"show,omitempty"`
	SkipTaskbar            *bool   `json:"skipTaskbar,omitempty"`
	Title                  *string `json:"title,omitempty"`
	Transparent            *bool   `json:"transparent,omitempty"`
	UseContentSize         *bool   `json:"useContentSize,omitempty"`
	Width                  *int    `json:"width,omitempty"`
	X                      *int    `json:"x,omitempty"`
	Y                      *int    `json:"y,omitempty"`
}

// NewWindow creates a new window
func (a *Astilectron) NewWindow(url string, o *WindowOptions) (w *Window, err error) {
	// Init
	w = &Window{
		c:  a.canceller,
		d:  a.dispatcher,
		id: a.identifier.new(),
		o:  o,
		w:  a.writer,
	}
	w.ctx, w.cancel = context.WithCancel(context.Background())

	// Make sure the window's context is cancelled once the closed event is received
	w.On(EventNameWindowEventClosed, func(e Event) (deleteListener bool) {
		w.cancel()
		return true
	})

	// Parse url
	if w.url, err = astiurl.Parse(url); err != nil {
		err = errors.Wrapf(err, "parsing url %s failed", url)
		return
	}
	return
}

// isActionable checks whether any type of action is allowed on the window
func (w *Window) isActionable() error {
	if w.isWindowDestroyed() {
		return errors.New("window has been destroyed")
	} else if w.c.Cancelled() {
		return errors.New("canceller has been cancelled")
	}
	return nil
}

// isWindowDestroyed checks whether the window has been destroyed
func (w *Window) isWindowDestroyed() bool {
	return w.ctx.Err() != nil
}

// On implements the Listenable interface
func (w *Window) On(eventName string, l Listener) {
	w.d.addListener(w.id, eventName, l)
}

// Blur blurs the window
func (w *Window) Blur() (err error) {
	if err = w.isActionable(); err != nil {
		return
	}
	return synchronousEvent(w.c, w, w.w, Event{Name: EventNameWindowCmdBlur, TargetID: w.id}, EventNameWindowEventBlur)
}

// Center centers the window
func (w *Window) Center() (err error) {
	if err = w.isActionable(); err != nil {
		return
	}
	return synchronousEvent(w.c, w, w.w, Event{Name: EventNameWindowCmdCenter, TargetID: w.id}, EventNameWindowEventMove)
}

// Close closes the window
func (w *Window) Close() (err error) {
	if err = w.isActionable(); err != nil {
		return
	}
	return synchronousEvent(w.c, w, w.w, Event{Name: EventNameWindowCmdClose, TargetID: w.id}, EventNameWindowEventClosed)
}

// CloseDevTools closes the dev tools
func (w *Window) CloseDevTools() (err error) {
	if err = w.isActionable(); err != nil {
		return
	}
	return w.w.write(Event{Name: EventNameWindowCmdWebContentsCloseDevTools, TargetID: w.id})
}

// Create creates the window
// We wait for EventNameWindowEventDidFinishLoad since we need the web content to be fully loaded before being able to
// send messages to it
func (w *Window) Create() (err error) {
	if err = w.isActionable(); err != nil {
		return
	}
	return synchronousEvent(w.c, w, w.w, Event{Name: EventNameWindowCmdCreate, TargetID: w.id, URL: w.url.String(), WindowOptions: w.o}, EventNameWindowEventDidFinishLoad)
}

// Destroy destroys the window
func (w *Window) Destroy() (err error) {
	if err = w.isActionable(); err != nil {
		return
	}
	return synchronousEvent(w.c, w, w.w, Event{Name: EventNameWindowCmdDestroy, TargetID: w.id}, EventNameWindowEventClosed)
}

// Focus focuses on the window
func (w *Window) Focus() (err error) {
	if err = w.isActionable(); err != nil {
		return
	}
	return synchronousEvent(w.c, w, w.w, Event{Name: EventNameWindowCmdFocus, TargetID: w.id}, EventNameWindowEventFocus)
}

// Hide hides the window
func (w *Window) Hide() (err error) {
	if err = w.isActionable(); err != nil {
		return
	}
	return synchronousEvent(w.c, w, w.w, Event{Name: EventNameWindowCmdHide, TargetID: w.id}, EventNameWindowEventHide)
}

// OpenDevTools opens the dev tools
func (w *Window) OpenDevTools() (err error) {
	if err = w.isActionable(); err != nil {
		return
	}
	return w.w.write(Event{Name: EventNameWindowCmdWebContentsOpenDevTools, TargetID: w.id})
}

// Maximize maximizes the window
func (w *Window) Maximize() (err error) {
	if err = w.isActionable(); err != nil {
		return
	}
	return synchronousEvent(w.c, w, w.w, Event{Name: EventNameWindowCmdMaximize, TargetID: w.id}, EventNameWindowEventMaximize)
}

// Minimize minimizes the window
func (w *Window) Minimize() (err error) {
	if err = w.isActionable(); err != nil {
		return
	}
	return synchronousEvent(w.c, w, w.w, Event{Name: EventNameWindowCmdMinimize, TargetID: w.id}, EventNameWindowEventMinimize)
}

// Move moves the window
func (w *Window) Move(x, y int) (err error) {
	if err = w.isActionable(); err != nil {
		return
	}
	w.o.X = PtrInt(x)
	w.o.Y = PtrInt(y)
	return synchronousEvent(w.c, w, w.w, Event{Name: EventNameWindowCmdMove, TargetID: w.id, WindowOptions: &WindowOptions{X: w.o.X, Y: w.o.Y}}, EventNameWindowEventMove)
}

// Resize resizes the window
func (w *Window) Resize(width, height int) (err error) {
	if err = w.isActionable(); err != nil {
		return
	}
	w.o.Height = PtrInt(height)
	w.o.Width = PtrInt(width)
	return synchronousEvent(w.c, w, w.w, Event{Name: EventNameWindowCmdResize, TargetID: w.id, WindowOptions: &WindowOptions{Height: w.o.Height, Width: w.o.Width}}, EventNameWindowEventResize)
}

// Restore restores the window
func (w *Window) Restore() (err error) {
	if err = w.isActionable(); err != nil {
		return
	}
	return synchronousEvent(w.c, w, w.w, Event{Name: EventNameWindowCmdRestore, TargetID: w.id}, EventNameWindowEventRestore)
}

// Send sends a message to the inner JS of the Web content of the window
func (w *Window) Send(message interface{}) (err error) {
	if err = w.isActionable(); err != nil {
		return
	}
	return w.w.write(Event{Message: newEventMessage(message), Name: EventNameWindowCmdMessage, TargetID: w.id})
}

// Show shows the window
func (w *Window) Show() (err error) {
	if err = w.isActionable(); err != nil {
		return
	}
	return synchronousEvent(w.c, w, w.w, Event{Name: EventNameWindowCmdShow, TargetID: w.id}, EventNameWindowEventShow)
}

// Unmaximize unmaximize the window
func (w *Window) Unmaximize() (err error) {
	if err = w.isActionable(); err != nil {
		return
	}
	return synchronousEvent(w.c, w, w.w, Event{Name: EventNameWindowCmdUnmaximize, TargetID: w.id}, EventNameWindowEventUnmaximize)
}
