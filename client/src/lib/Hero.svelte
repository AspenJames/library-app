<script>
    import Button from "./Button.svelte";
    import Card from "./Card.svelte";
</script>

<div class="container">
    <div class="card">
        <Card>
            <p>Create a digital inventory of your physical books</p>
            <div class="buttons">
                <Button button_text={"Register"} orange_button></Button>
                <Button button_text={"Log In"}></Button>
            </div>
        </Card>
   </div>
</div>
 
<style>
    .container{
        background-image: url("./src/assets/homepage_hero_image.png");
        background-size: cover;
        width: 100%;
        height: 100%;
        display: grid;
        grid-template-columns: 1280px;
        grid-template-rows: repeat(3, 1fr);
    }
    .card{
        grid-column-start:1;
        grid-row-start:2;
        grid-row-end:3;
        justify-self: center;
        display: grid;
        grid-template-columns: auto;
        grid-template-rows: repeat(2, 50%);

        width: 600px;
        height: 200px;
        box-sizing: border-box;
        background: white;
        border-radius: 30px;
        color: rgb(75, 75, 75);
    }
    .card p{
        line-height: 1.2;
        text-align: center; 
        vertical-align: middle;
        display: table-cell;
        font-size: 40px;
        font-weight: bold;
        color: black;
    }
    .buttons{
        grid-column-start:1;
        grid-row-start:2;
        grid-row-end:3;
        justify-self: center;
        align-self: center;
        margin-top: 20px;
        }
</style>