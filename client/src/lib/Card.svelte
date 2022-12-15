<!--card component-->
<script>
    export let homepage_card=false;
</script>

<div class="card" class:homepage_card><slot /></div>

<!--add buttons to cards here-->

<style>
    .card{
        --width: 400px;
        --height: 320px;
        --box-shadow: 0 0 200px #0080ff;
        --opacity: 1;

        width: var(--width);
        height: var(--height);
        box-shadow: var(--box-shadow);
        opacity: var(--opacity);
        box-sizing: border-box;
        background: white;
        border-radius: 30px;
        color: rgb(75, 75, 75);
    }
    .card.homepage_card{
        --width: 600px;
        --height: 200px;
        --box-shadow: none;
    }
</style>