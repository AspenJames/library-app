<script>
    import { createEventDispatcher } from 'svelte';

    const dispatch = createEventDispatcher();
    export let filterAll=true;
    function eventFilterAllClicked() {
        dispatch('filter-all');
        };
</script>

{#if filterAll}
    <button class=active on:click|preventDefault={eventFilterAllClicked}>All</button>
{:else}
    <button class=inactive on:click|preventDefault={eventFilterAllClicked}>All</button>
{/if}

<style>
    button{
        --border-bottom-width: 4px;
        --border-bottom-color: white;
        --color: #555555;   
        --box-shadow: none;
        --font-weight: normal;

        border-bottom-width: var(--border-bottom-width);
        border-bottom-color: var(--border-bottom-color);
        color: var(--color);
        background-color: white;
        border-radius: 0%;
        font-weight: var(--font-weight);
        font-size: 14px;
        height: auto;
        width: auto;
        letter-spacing: -0.02em;
        padding-bottom: 0px;
        padding-top: 2px;
    }
    button.active{
        border-bottom: 5px black solid;
        font-weight: bold;
        color: black;
        background-color: rgb(214, 214, 214);
    }
    button.inactive{
        border-bottom: 5px rgb(235,235,235) solid;
        font-weight: normal;
        color: grey;
        background-color: white;
}
</style>