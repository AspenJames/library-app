<script>
    export let large=false;
</script>

<img src='./src/assets/logo.png' class="logo" class:large alt="Library Logo"/>

<style>
    .logo{
        --height: auto;
        --width: auto;

        height: var(--height);
        width: var(--width);
        padding: 0;
        will-change: filter;
        max-width: 100%;
        max-height: 100%;
    }
    .logo.large{
        --height: 30%;
        --width: 30%;
    }
    .logo.large:hover{
        filter: drop-shadow(0 0 40px #005fdc);
    }
</style>