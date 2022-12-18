<!--library filter menu component-->
<script>
    export let active=true;
</script>

<div class="buttons">
    <button class="menu_button" class:active>All</button>
    <button class="menu_button">Unread</button>
    <button class="menu_button">Read</button>
</div>


<style>
    .buttons{
        display: flex;
        flex-direction: row ;
        width: 150px;
    }
    .menu_button{
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
    .menu_button.active{
        --border-bottom-width: 4px;
        --border-bottom-color: #FF8A00;
        --color: #FF8A00;
        --font-weight: 500;
    }
</style>