# Navigating between routes

In Angular, we can make our app move between different pages (or "routes") without the user needing
to change the URL manually. There are two ways to do this:

`routerLink` is like a clickable link that takes the user to a new page
 For example, if we want to add a link to the settings page, we use this in the template:

```html
<a routerLink="/settings">Go to settings</a>
```

This makes a simple link that, when clicked, takes the user to the "settings" page.

`router` if we need to do something first (like logging in) before navigating.
We can use `router.navigate()` in the component class programmatically.

For example:

```ts
router = inject(Router);

handleLogin() {
    this.router.navigate(['settings']);
}
```

This lets us control when and how to go to the next page based on actions like button clicks.

We can also use **relative paths** (e.g., `./settings`) if we want the route to be based on where
we currently are in the app.

```ts
router = inject(Router);
route = inject(ActivatedRoute); // ActivatedRoute gets information about the current route

handleLogin() {
    this.router.navigate(['settings'], { relativeTo: this.route });
}
```

## Using route params

Sometimes when navigating to a page, we want to send some information along with the navigation.
This is where route parameters come in.

For example, if we want to go to a page that shows details about a movie, we can pass the movie's ID
in the URL like this:

```ts
{ path: 'movie/:id', ... } // :id is a placeholder
```

When navigating, we can use and tells Angular to go to the "movie" page with the ID of 1

```html
<a routerLink="/movie/1">Dune</a>
or
<a [routerLink]="['/movie', movie.id]">Dune</a>
```

or

```ts
this.router.navigate(['movie', this.movie.id]);
```

## Passing data between pages

When building apps, we often want to pass information between pages.
For example, on the home page, we might show a list of movies, and when we click on a movie,
we want to go to a detail page with more information about that movie.

We first set up routes:

```ts
{
    path: 'home',
    loadComponent: () => import('./home/home.component').then((m) => m.HomeComponent),
},
{
    path: 'detail/:id',
    loadComponent: () => import('./detail/detail.component').then((m) => m.DetailComponent),
},
```

On the home page, we can list the movies and use `routerLink` to pass the movie ID to the detail page.

```html
<a [routerLink]="['/detail', movie.id]">{{ movie.name }}</a>
```

Then we fetch data on detail page On the detail page, we can use the movie ID from the route to
fetch the full movie data from a service

```ts
// or with the withComponentInputBinding() options we can use the input signal
id = input();

ngOnInit() {
    const id = this.route.snapshot.params['id'];
    this.moviesService.getMovieById(id).subscribe((movie) => {
        this.movie = movie;
    });
}
```
