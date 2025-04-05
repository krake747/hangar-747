# State Management

State is what happens when something in our app can change.
For example, a lightbulb can be on or off, and an elevator can be at different levels.
In our app, state is what determines if something can change based on user actions.

Initially, our app might not have state, like when we just display a greeting.
But once we add a button to toggle the greeting visibility, we introduce state.
The button changes the value of `showGreeting`, making the app reactively show or hide the greeting.

Managing state can be simple at first, but as apps grow, it can get complicated.
Many strategies exist for state management, and people have different opinions on the best approach
and libraries to use.
