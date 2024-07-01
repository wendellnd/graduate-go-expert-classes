```
mutation createCourse {
  createCourse(input: {name: "Full Cycle", description: "Tecnologia", categoryId: "354332e2-e5e6-47cd-89a4-049456f5c6d3"}) {
    id,
    name,
    description,
    Category{
      id,
    }
  }
}

query queryCourses {
  courses{
    id,
    name,
    description
	}
}

query queryCoursesWithCourses {
  categories{
    id,
    name,
    description,
    courses{
			id,
      name
    }
	}
}

query queryCoursesWithCategory {
  courses{
    id,
    name,
    description,
    Category{
      id,
      name,
      description
    }
  }
}
```
