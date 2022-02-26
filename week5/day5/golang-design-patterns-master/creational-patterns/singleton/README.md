#### Introduction

Sometimes it's important to have only one instance of an struct. This is useful
when exactly one object is needed to coordinate actions across the system.
Singletons provide a global point of access to themselves.

The singleton pattern is one of the simplest design patterns. It requires only
one type which is responsible to instantiate itself, to make sure it creates
not more than one instance. The same instance can be used from everywhere.

#### Purpose

- Ensure that only one instance of a class is created for application lifecycle.
- Provide a single point of access to the object.

#### Desing Pattern Diagram

![alt tag](http://blog.ralch.com/media/golang/design-patterns/singleton.gif)

#### Implementation

In Golang the Singleon pattern implementation defers from another languages due 
to the language differences.

In langauges like `C#` and `JAVA`, the implementation involves a static member
in the `Singleton` class, a private constructor and a static public method that
returns a reference to the static member.

In Golang the static member of the `Singleton` struct is declared as a global
variable in the package that contains this type.

Lets have a `db` package that should provide a `repository` struct as a
singleton object. Note that we should define the struct with lowercase letters
in order to make it private. This will disallow instaciating the struct outside
the package.

```Golang
package db

import "fmt"

type repository struct {
	items map[string]string
	mu    sync.RWMutex
}

func (r *repository) Set(key, data string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.items[key] = data
}

func (r *repository) Get(key string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	item, ok := r.items[key]
	if !ok {
		return "", fmt.Errorf("The '%s' is not presented", key)
	}
	return item, nil
}

```

Then we should declare a global variable of type pointer to `repository` that will
keep an reference to its singleton instance. Then we should declare a function
`Repository` that provides a global point of access to that instance.

```Golang
var (
	r *repository
)

func Repository() *repository {
	if r == nil {
		r = &repository{
			items: make(map[string]string)
		}
	}
	return r
}
```

The function `Repository` instanciate the singleton object once if it has not
been instance. It checks whether the `repository` global variable is `nil`.

#### Thread safety

A robust singleton implementation should work in any circumstances. This is why
we need to ensure it works when multiple go routines use it.

The standard implementation requires to synchronize the action that instanciate
the singleton object once.

In order to achieve that we should use the [sync](https://golang.org/pkg/sync/)
package. It provides a [sync.Once](https://golang.org/pkg/sync/#Once) struct
that will perform exactly one action:

```Golang
var (
	r    *repository
	once sync.Once
)

func Repository() *repository {
	once.Do(func() {
		r = &repository{
			items: make(map[string]string),
		}
	})
	
	return r
}
```

#### Verdict

The Singleton design pattern is a very useful mechanism for providing a single
point of object access in an object-oriented application. Regardless of the
implementation used, the pattern provides a commonly understood concept that
can be easily shared among design and development teams.
