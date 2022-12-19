<script>
    export let homepage=false;
    import Button from './Button.svelte';
    import Logo from './Logo.svelte';
</script>

<div class=customHeader>
{#if homepage==false}
    <div class="logo">
        <Logo></Logo>
    </div>
    <div class="logOutButton">
        <Button button_text="Log Out"></Button>
    </div>
{:else}
    <div class="logo">
        <Logo></Logo>
    </div>
{/if}
</div>

<style>
    .customHeader{
        --header-height: 50px;
        --header-width: 1280px;
        --horizontal-margins-width: 100px;

        height: var(--header-height);
        width: var(--header-width);
        max-width: 100%;
        box-shadow: 0px 2px 7px rgb(113, 113, 112);
        background-color: white;
        display: grid;
        grid-template-columns: 1fr 4fr;
        grid-template-rows: var(--header-height);
        margin-bottom: 10px;
        margin-left: var(--horizontal-margins-width);
        margin-right: var(--horizontal-margins-width);
    }
    .logo{
        grid-row-start: 1;
        grid-column-start: 1;
        grid-column-end: 2;
        justify-self: start;
        margin-left: 10px;
    }
    .logOutButton{
        grid-row-start: 1;
        grid-column-start: 2;
        grid-column-end: 3;
        justify-self: end;
        align-self: center;
        margin-right: 10px;
    }
</style>