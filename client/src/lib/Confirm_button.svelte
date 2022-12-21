<script>
    import { createEventDispatcher } from 'svelte';

    const dispatch = createEventDispatcher();

    function eventConfirmButtonClicked(confirm_delete_event) {
        dispatch('message', {text:'confirm button clicked'}
        )};
</script>

<button on:click|preventDefault={eventConfirmButtonClicked}>CONFIRM</button>

<style>
    button {
        color: red;
        background-color: white;
		width: 48px;
		font-size: 9px;
		height: 15px;
		border-radius: 2px;
		padding: 0px;
		text-align: center;
		font-weight: bold;
        will-change: filter;
        border: 2px solid red;
        }
    button:hover {
        background-color: lightpink;
        width: 65px;
        height: 20px;
        font-size: 12px;
        text-shadow: 0px 0px 5px #ffbebe;
        color: rgb(189, 0, 0);
        border: 4px solid rgb(189, 0, 0);
    }
</style>