<!--library filter menu component-->
<script>
    export let active=false;
</script>

<button class="menu_button" class:active>All</button>
<button class="menu_button">Unread</button>
<button class="menu_button">Read</button>


<style>
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
        height: 50px;
        width: 160px;
        letter-spacing: -0.02em;
        float: left;
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