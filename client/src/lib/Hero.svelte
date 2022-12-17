<script>
    import Button from "./Button.svelte";
    import Card from "./Card.svelte";
    import { slide , fade } from 'svelte/transition';
    let cardVisible=true;
    let cardRegisterVisible=false;
    let cardLogInVisible=false;
    let name = 'username';
    let password = 'password';

    function handleRegisterClick() {
        cardVisible=false;
        cardRegisterVisible=true;
        cardVisible=cardVisible;
        cardRegisterVisible=cardRegisterVisible;
    }

    function handleLogInClick() {
		cardVisible=false;
        cardLogInVisible=true;
        cardVisible=cardVisible;
        cardLogInVisible=cardLogInVisible;
    };
    function handleFinalRegisterClick() {
        alert('You are registered. Taking you to library page.')
    }
    function handleFinalLogInClick() {
        alert('You are logged in. Taking you to library page.')
    }
</script>

<div class="container"> 
    {#if cardVisible}
        <div class="card" in:slide="{{ duration: 1000 }}" out:slide="{{ duration: 1000 }}">
            <div class=cardP>
                <p>Create a digital inventory of your physical books</p>
            </div>
            <div class="buttons" style="margin-top: 20px;">
                <Button button_text={"Register"} orange_button on:click={handleRegisterClick}></Button>
                <Button button_text={"Log In"} on:click={handleLogInClick}></Button>
            </div>
        </div>
    {:else if cardRegisterVisible}
        <div class="cardRegister" in:slide="{{ duration: 1000 }}" out:slide="{{ duration: 1000 }}">
            <div class=cardRegisterP>
                <p>Register for an account</p>
            </div>
            <div class="inputs">
                <input bind:value={name}>
                <input bind:value={password}>
            </div>
            <div class="buttonsRegister">
                <Button button_text={"Register"} wide_orange_button on:click={handleFinalRegisterClick}></Button>
            </div>
        </div>
    {:else if cardLogInVisible}
        <div class="cardRegister" in:slide="{{ duration: 1000 }}" out:slide="{{ duration: 1000 }}">
            <div class=cardRegisterP>
                <p>Log in to your account</p>
            </div>
            <div class="inputs">
                <input bind:value={name}>
                <input bind:value={password}>
            </div>
            <div class="buttonsRegister">
                <Button button_text={"Log In"} wide_orange_button on:click={handleFinalLogInClick}></Button>
            </div>
        </div>
    {/if}
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
    .cardRegister{
        grid-column-start:1;
        grid-row-start:2;
        grid-row-end:3;
        justify-self: center;
        display: flex;
        flex-direction: column;
        justify-content: flex-start;
        align-content: space-between;
        width: 600px;
        height: 200px;
        box-sizing: border-box;
        background: white;
        border-radius: 30px;
        color: rgb(75, 75, 75);
    }
    .cardP{
        line-height: 1.1;
        font-size: 40px;
        font-weight: bold;
        color: black;

        grid-column-start:1;
        grid-row-start:1;
        grid-row-end:2;
        align-self: center;
    }
    .cardRegisterP{
        line-height: 0.5;
        text-align: center; 
        vertical-align: top;
        font-size: 40px;
        font-weight: bold;
        color: black;
        order: 1;
    }
    input{
        background-color: white;
        border-top-color: none;
        border-color: grey;
        border-width: 0.5px;
        color: grey;
        margin-bottom: 20px;
        margin-left: 10px;
        margin-right: 10px;
    }
    .inputs{
        order: 2;
    }
    .buttons{
        grid-column-start:1;
        grid-row-start:2;
        grid-row-end:3;
        justify-self: center;
        align-self: center;
        }
    .buttonsRegister{
        order: 3;
        justify-self: center;
        align-self: center;
        }
</style>