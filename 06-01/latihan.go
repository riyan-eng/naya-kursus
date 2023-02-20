package main

import "fmt"

type SosmedInterface interface {
	Status() error
	Share() error
	Message() error
}

type Instagram struct {
	Name string
}

type TikTok struct {
	Name string
}

type WhatsApp struct {
	Name string
}

func NewInstagram(name string) SosmedInterface {
	return &Instagram{
		Name: name,
	}
}

func (ig Instagram) Status() error {
	fmt.Println("status", ig.Name)
	return nil
}

func (ig Instagram) Share() error {
	fmt.Println("share", ig.Name)
	return nil
}

func (ig Instagram) Message() error {
	fmt.Println("message", ig.Name)
	return nil
}

func NewTikTok(name string) SosmedInterface {
	return &TikTok{
		Name: name,
	}
}

func (tt TikTok) Status() error {
	fmt.Println("status", tt.Name)
	return nil
}

func (tt TikTok) Share() error {
	fmt.Println("share", tt.Name)
	return nil
}

func (tt TikTok) Message() error {
	fmt.Println("message", tt.Name)
	return nil
}

func NewWhatsApp(name string) *WhatsApp {
	return &WhatsApp{
		Name: name,
	}
}

func (wa WhatsApp) Status() error {
	fmt.Println("status", wa.Name)
	return nil
}

func (wa WhatsApp) Share() error {
	fmt.Println("share", wa.Name)
	return nil
}

func (wa WhatsApp) Message() error {
	fmt.Println("message", wa.Name)
	return nil
}

func (wa WhatsApp) Call() error {
	fmt.Println("call", wa.Name)
	return nil
}

func (wa WhatsApp) VidCall() error {
	fmt.Println("vidcall", wa.Name)
	return nil
}

func main() {
	// ig
	ig := NewInstagram("ig")
	ig.Status()
	ig.Message()
	ig.Share()

	fmt.Println()

	// tt
	tt := NewTikTok("tt")
	tt.Message()
	tt.Share()
	tt.Status()

	fmt.Println()

	// wa
	wa := NewWhatsApp("wa")
	wa.Message()
	wa.Share()
	wa.Status()
	wa.Call()
	wa.VidCall()
}
