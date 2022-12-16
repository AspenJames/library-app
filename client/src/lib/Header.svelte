<script>
    export let homepage=false;
    import Button from './Button.svelte';
    import Logo from './Logo.svelte';
</script>
{#if homepage=false}
    <div class="logo">
        <Logo></Logo>
    </div>
    <div class="buttons">
        <Button></Button>
        <Button></Button>
    </div>
{:else}
    <div class="logo">
        <Logo></Logo>
    </div>
{/if}
<style>
    .logo{
        --header-height: 50px;
        --header-width: 1280px;
        --horizontal-margins-width: 100px;

        height: var(--header-height);
        width: var(--header-width);
        max-width: 100%;
        box-shadow: 0px 3px 4px rgb(47, 47, 47);
        background-color: white;
        display: grid;
        grid-template-columns: 200px auto 150px 150px; /*770 was found through trial and error. need to clean this up*/
        grid-template-rows: var(--header-height);
        margin-bottom: 10px;
        margin-left: var(--horizontal-margins-width);
        margin-right: var(--horizontal-margins-width);
    }
</style>