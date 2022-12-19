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
		width: 45px;
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
        filter: drop-shadow(1px 1px 10px red);
        background-color: lightpink;
        width: 58px;
        height: 20px;
        font-size: 12px;
        text-shadow: 0px 0px 5px #ffbebe;
    }
</style>