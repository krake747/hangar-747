# Why a state management approach matters

In Angular, we might hear that we don't need complex state management.
A common approach is using a service with a `signal` (or `Subject`), like this:

```ts
@Injectable({
     providedIn: 'root',
})
export class PreferenceService {
  #showGreeting = signal(false);
  showGreeting = this.#showGreeting.asReadonly();

  toggleGreeting() {
    this.#showGreeting.update((showGreeting) => !showGreeting);
  }
}
```

This approach seems simple and fine, but the problem is we're still managing state,
and if we don't have a clear strategy for it, our approach becomes messy.
Without a defined state management strategy, we could end up with inconsistent solutions that confuse
our team.
Even experienced developers might use an ad-hoc approach, but this doesn't work well in team environments.

It's better to choose a clear, consistent state management strategy so we can communicate better
with other developers and maintain a clean codebase.
