package builder

// Product represents the complex object being built
type Product struct {
    partA string
    partB string
    partC string
}

// GetPartA returns part A of the product
func (p *Product) GetPartA() string {
    return p.partA
}

// GetPartB returns part B of the product
func (p *Product) GetPartB() string {
    return p.partB
}

// GetPartC returns part C of the product
func (p *Product) GetPartC() string {
    return p.partC
}

// Builder is the interface that specifies methods for creating parts of the Product
type Builder interface {
    BuildPartA()
    BuildPartB()
    BuildPartC()
    GetProduct() *Product
}

// ConcreteBuilder1 implements the Builder interface for the first type of product
type ConcreteBuilder1 struct {
    product *Product
}

// NewConcreteBuilder1 creates a new ConcreteBuilder1
func NewConcreteBuilder1() *ConcreteBuilder1 {
    return &ConcreteBuilder1{
        product: &Product{},
    }
}

// BuildPartA builds part A of the product
func (b *ConcreteBuilder1) BuildPartA() {
    b.product.partA = "PartA1"
}

// BuildPartB builds part B of the product
func (b *ConcreteBuilder1) BuildPartB() {
    b.product.partB = "PartB1"
}

// BuildPartC builds part C of the product
func (b *ConcreteBuilder1) BuildPartC() {
    b.product.partC = "PartC1"
}

// GetProduct returns the built product
func (b *ConcreteBuilder1) GetProduct() *Product {
    return b.product
}

// ConcreteBuilder2 implements the Builder interface for the second type of product
type ConcreteBuilder2 struct {
    product *Product
}

// NewConcreteBuilder2 creates a new ConcreteBuilder2
func NewConcreteBuilder2() *ConcreteBuilder2 {
    return &ConcreteBuilder2{
        product: &Product{},
    }
}

// BuildPartA builds part A of the product
func (b *ConcreteBuilder2) BuildPartA() {
    b.product.partA = "PartA2"
}

// BuildPartB builds part B of the product
func (b *ConcreteBuilder2) BuildPartB() {
    b.product.partB = "PartB2"
}

// BuildPartC builds part C of the product
func (b *ConcreteBuilder2) BuildPartC() {
    b.product.partC = "PartC2"
}

// GetProduct returns the built product
func (b *ConcreteBuilder2) GetProduct() *Product {
    return b.product
}

// Director constructs the product using the Builder interface
type Director struct {
    builder Builder
}

// NewDirector creates a new Director with the specified builder
func NewDirector(builder Builder) *Director {
    return &Director{
        builder: builder,
    }
}

// Construct builds the product by calling the builder's methods in a specific order
func (d *Director) Construct() *Product {
    d.builder.BuildPartA()
    d.builder.BuildPartB()
    d.builder.BuildPartC()
    return d.builder.GetProduct()
}

// SetBuilder allows changing the builder at runtime
func (d *Director) SetBuilder(builder Builder) {
    d.builder = builder
} 