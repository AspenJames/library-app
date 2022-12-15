<script>
    import Button from "./Button.svelte";
    import Card from "./Card.svelte";
</script>

<!--usebackground-image property in div instead of a div in a div - then only use one div for image + card overlay-->

<div class="container">
    <div class="item_transparent_card"> <!--since this div only contains card this is not acting on the item button. need to lay button right in card. do this in card component-->
        <Card homepage_card>
            <p>Create a digital inventory of your physical books</p>
            <div class="item_button">
                <Button orange_button button_text={"Sign Up"}></Button>
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
    .item_transparent_card{
        grid-column-start:1;
        grid-row-start:2;
        grid-row-end:3;
        justify-self: center;
        display: grid;
        grid-template-columns: auto;
        grid-template-rows: repeat(2, 50%);
    }
    .item_transparent_card p{
        line-height: 1.2;
        text-align: center; 
        vertical-align: middle;
        display: table-cell;
        font-size: 40px;
        font-weight: bold;
        color: black;
    }
    .item_button{
        grid-column-start:1;
        grid-row-start:2;
        grid-row-end:3;
        justify-self: center;
        align-self: center;
        opacity: 1;
        margin-top: 20px;
        }
</style>