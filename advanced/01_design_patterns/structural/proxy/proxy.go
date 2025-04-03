package proxy

import (
	"fmt"
	"sync"
	"time"
)

// Subject defines the common interface for RealSubject and Proxy
type Subject interface {
    Request() string
}

// RealSubject represents the real object that the proxy represents
type RealSubject struct {
    name string
}

// NewRealSubject creates a new RealSubject
func NewRealSubject(name string) *RealSubject {
    return &RealSubject{name: name}
}

// Request implements the Subject interface
func (s *RealSubject) Request() string {
    return "RealSubject " + s.name + " request"
}

// Proxy controls access to the RealSubject
type Proxy struct {
    realSubject *RealSubject
    mutex       sync.Mutex
    cache       string
    lastAccess  time.Time
    cacheTTL    time.Duration
}

// NewProxy creates a new Proxy
func NewProxy(realSubject *RealSubject, cacheTTL time.Duration) *Proxy {
    return &Proxy{
        realSubject: realSubject,
        cacheTTL:    cacheTTL,
    }
}

// Request implements the Subject interface
func (p *Proxy) Request() string {
    p.mutex.Lock()
    defer p.mutex.Unlock()

    // Check if cache is valid
    if !p.lastAccess.IsZero() && time.Since(p.lastAccess) < p.cacheTTL {
        return p.cache
    }

    // Perform the real request
    result := p.realSubject.Request()
    
    // Update cache
    p.cache = result
    p.lastAccess = time.Now()
    
    return result
}

// ProtectionProxy controls access to the RealSubject based on permissions
type ProtectionProxy struct {
    realSubject *RealSubject
    permission  string
}

// NewProtectionProxy creates a new ProtectionProxy
func NewProtectionProxy(realSubject *RealSubject, permission string) *ProtectionProxy {
    return &ProtectionProxy{
        realSubject: realSubject,
        permission:  permission,
    }
}

// Request implements the Subject interface
func (p *ProtectionProxy) Request() string {
    if p.permission != "admin" {
        return "Access denied"
    }
    return p.realSubject.Request()
}

// VirtualProxy controls access to a resource that is expensive to create
type VirtualProxy struct {
    realSubject *RealSubject
    initialized bool
}

// NewVirtualProxy creates a new VirtualProxy
func NewVirtualProxy() *VirtualProxy {
    return &VirtualProxy{
        initialized: false,
    }
}

// Request implements the Subject interface
func (p *VirtualProxy) Request() string {
    if !p.initialized {
        fmt.Println("Initializing RealSubject...")
        p.realSubject = NewRealSubject("virtual")
        p.initialized = true
    }
    return p.realSubject.Request()
}

// Client represents a client that uses the Subject interface
type Client struct {
    subject Subject
}

// NewClient creates a new Client
func NewClient(subject Subject) *Client {
    return &Client{subject: subject}
}

// UseSubject uses the Subject interface
func (c *Client) UseSubject() string {
    return c.subject.Request()
} 