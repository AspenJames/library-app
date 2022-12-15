<script>
    export let with_buttons=false;
    import Button from './Button.svelte';
    import Logo from './Logo.svelte'
</script>

<div class="container" class:with_buttons>
    <div class="item-a">
        <Logo></Logo>
    </div>
    <div class="item-b">
        <Button button_text={'Log In'}></Button>
    </div>
    <div class="item-c">
        <Button orange_button button_text={'Sign Up'}></Button>
    </div>
</div>

<style>
    .container{
        --header-height: 50px;
        --header-width: 1280px;

        height: var(--header-height);
        width: var(--header-width);
        box-shadow: 0px 3px 4px rgb(47, 47, 47);
        background-color: white;
        display: grid;
        grid-template-columns: 200px 740px 150px 150px; /*770 was found through trial and error. need to clean this up*/
        grid-template-rows: var(--header-height);
        margin-bottom: 10px;
    }
    .item-a{
        grid-column-start: 1;
        grid-row-start: 1;
        margin: 5px;
        justify-self: start;
    }
    .item-b{
        grid-column-start: 3;
        grid-row-start: 1;
        justify-self: end;
        align-self: center;
    }
    .item-c{
        grid-column-start: 4;
        grid-row-start: 1;
        justify-self: end;
        align-self: center;
    }
</style>