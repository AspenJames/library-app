<script>
    import { createEventDispatcher } from 'svelte';

    const dispatch = createEventDispatcher();

    function eventDeleteButtonClicked(delete_event) {
        dispatch('message', {text:'del button clicked'}
        )};
</script>

<button on:click|preventDefault={eventDeleteButtonClicked}>X</button>

<style>
    button {
        background-color: inherit;
        border-radius: 0;
        color: red;
        padding: 0;
        text-align: center;
        text-decoration: none;
        font-size: 16px;
        will-change: filter;
        }
    button:hover {
        filter: drop-shadow(1px 1px 5px red);
        font-weight: bold;
        background-color: lightpink;
    }
</style>