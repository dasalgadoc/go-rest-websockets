# 📑🐹 Pagination and custom Queries

What happens if you only want a select few elements of your persistence?

There are two ways to answer this question:

## 👶🏻 A naive implementation

Use SQL (or any DQL implementation) to reduce the number of results returned from a query:

Postgres

```postgresql
-- Pages 10 to 10 (10 to 20 elements)
SELECT *
FROM PRODUCTS
WHERE p.Category
LIMIT 10 OFFSET 10:
```

SQL Server

```tsql
-- Pages 10 to 10 (10 to 20 elements)
SELECT TOP 10 *
FROM PRODUCTS AS p WITH (NOLOCK)
WHERE p.Category = $1
ORDER BY Name
OFFSET 10 ROWS;
```

When solving custom procedures, SQL offers a variety of options, combining  _GROUP BY_, _HAVING_, _SORT_, _ROW_COUNT, etc.

But, exists an alternative.

## 🦸🏻‍ Specification Pattern (aka criteria)

[Read more](https://deviq.com/design-patterns/specification-pattern)
[Read more](https://medium.com/c-sharp-progarmming/specification-design-pattern-c814649be0ef)

Let Products be a struct

```go
type Product struct {
    Name string
    Price float64
    Category string
}
```

Create an interface named "_Specification_" if you wish to search products based on a list of parameters like name, price, or category.

```go
type Specification interface {
	IsSatisfiedBy(product *Product) bool
}
```

Each implementation define a search criteria

```go
type NameSpecification struct {
	name string
}

func (s * NameSpecification) IsSatisfiedBy(product *Product) bool {
	return product.Name = s.name
}
```

In the repository generate the function _FindAll_

```go
type ProductRepository struct {
	products []*Product
} 

func (r *ProductRepository) FindAll(spec Specification) []*Product {
   var result []*Product
   for _, product := range r.products {
       if spec.IsSatisfiedBy(product) {
           result = append(result, product)
       }
   }
   return result
}

// Example
products := []*Product{
    &Product{Name: "iPhone", Price: 999.99, Category: "Electronics"},
    &Product{Name: "MacBook Pro", Price: 1999.99, Category: "Computers"},
    &Product{Name: "AirPods", Price: 149.99, Category: "Electronics"},
}

repo := &ProductRepository{products}

// ByName
iphoneSpec := &NameSpecification{name: "iPhone"}
iPhones := repo.FindAll(iphoneSpec)

// ByCategory
electronicsSpec := &CategorySpecification{category: "Electronics"}
electronics := repo.FindAll(electronicsSpec)

// ByPrice
expensiveSpec := &PriceSpecification{minPrice: 1000.0}
expensiveProducts := repo.FindAll(expensiveSpec)
```
