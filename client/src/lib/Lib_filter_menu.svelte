<!--library filter menu component-->

<!--export an event dispatcher to the lib_page to use button clicks in if-else blocks to get filtered results-->
<script>
    export let activeAll=true;
    export let activeRead=false;
    export let activeUnread=false;

    function filterAll() {
        activeAll=true;
    }
    
</script>

<button class:active={true} on:click><slot /></button>

<style>
    button{
        --border-bottom-width: 4px;
        --border-bottom-color: white;
        --color: #555555;
        --box-shadow: none;
        --font-weight: 200;

        border-bottom-width: var(--border-bottom-width);
        border-bottom-color: var(--border-bottom-color);
        color: var(--color);
        background-color: white;
        border-radius: 0%;
        font-weight: var(--font-weight);
        font-size: 16px;
        height: auto;
        width: auto;
        letter-spacing: -0.02em;
        padding-bottom: 0px;
        padding-top: 2px;
        will-change: filter;
    }
    button.active{
        border-bottom: 5px orange solid;
        font-weight: bold;
        color: orange;
        background-color: rgb(253, 236, 205);
    }
</style>