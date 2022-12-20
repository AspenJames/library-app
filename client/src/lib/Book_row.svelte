<!--book row component-->
<script>
    import { fly } from 'svelte/transition';
    import { onMount } from 'svelte';
    import RowTag from './Row_tag.svelte';
    import DeleteButton from './Delete_button.svelte';
    import ConfirmButton from './Confirm_button.svelte';
    import { backIn, elasticIn, expoOut, quintIn, quintOut, cubicOut, linear, quartOut, quadIn, sineIn, backOut, elasticOut, bounceOut } from 'svelte/easing';
    export let title=undefined;
    export let author=undefined;
    export let edition=undefined;
    export let ISBN=undefined;
    export let read_status=false;

    let transitionOnPageLoad = false;
    onMount(() => transitionOnPageLoad = true);

    let normalRowSize = true;
    let randomY = 1;
    let randomX = 1;
    let randomDelay = 50;
    let randomDuration = 50;

    function handleConfirmDeleteBookClick(confirm_delete_event) {
        normalRowSize=true;
        normalRowSize=normalRowSize;
    }
    //generates random integer between min and max inclusively
    function getRndInteger(min=-250, max=250) {
        return Math.floor(Math.random() * (max - min + 1) ) + min;
    }
    function genRandomX_Y () {{
        randomY = getRndInteger();
        randomY=randomY;
        randomX = getRndInteger();
        randomX=randomX;
    }  }
    function genRandomDelay () {
        randomDelay = getRndInteger(250, 350);
        randomDelay=randomDelay;
    }
    function genRandomDuration () {
        randomDuration = getRndInteger(1000, 1500);
        randomDuration=randomDuration;
    }
    function handleDeleteBookClick(delete_event) {
        normalRowSize=false;
        normalRowSize=normalRowSize;
    };



    genRandomX_Y();
    genRandomDelay();
    genRandomDuration();
   
/*
    export interface FlyParams {
	delay?: number;
	duration?: number;
	easing?: EasingFunction;
	x?: number;
	y?: number;
	opacity?: number;
}

export function fly(node: Element, {
	delay = 0,
	duration = 400,
	easing = cubicOut,
	x = 0,
	y = 0,
	opacity = 0
}: FlyParams = {}): TransitionConfig {
	const style = getComputedStyle(node);
	const target_opacity = +style.opacity;
	const transform = style.transform === 'none' ? '' : style.transform;

	const od = target_opacity * (1 - opacity);

	return {
		delay,
		duration,
		easing,
		css: (t, u) => `
			transform: ${transform} translate(${(1 - t) * x}px, ${(1 - t) * y}px);
			opacity: ${target_opacity - (od * u)}`
	};
}

   */ 
</script>
{#if transitionOnPageLoad}
    {#if normalRowSize}
        <div class="book_row" in:fly="{{ x: randomX, y:randomY, delay: randomDelay, duration: randomDuration, easing: bounceOut }}">
            <div class="item-tag">
                <RowTag read_status={read_status}></RowTag>
            </div>
            <div class="item-title">
                <p>{title}</p>
            </div>
            <div class="item-author">
                <p>{author}</p>
            </div>
            <div class="item-edition">
                <p>{edition}</p>
            </div>
            <div class="item-isbn">
                <p>{ISBN}</p>
            </div>
            <!--<DeleteButton on:message={forward}></DeleteButton> -->
            <DeleteButton on:message={handleDeleteBookClick}></DeleteButton>
        </div>
    {:else}
        <div class="book_row_short" out:fly="{{ x:-500, delay: 200, duration: 800, easing: backIn }}">
            <div class="item-tag">
                <RowTag read_status={read_status}></RowTag>
            </div>
            <div class="item-title">
                <p>{title}</p>
            </div>
            <div class="item-author">
                <p>{author}</p>
            </div>
            <div class="item-edition">
                <p>{edition}</p>
            </div>
            <div class="item-isbn">
                <p>{ISBN}</p>
            </div>
            <!--<DeleteButton on:message={forward}></DeleteButton> -->
            <ConfirmButton on:message={handleConfirmDeleteBookClick} on:message={genRandomX_Y}>CONFIRM</ConfirmButton>
        </div>
    {/if}
{/if}


<style>
    .book_row{
        --custom-grid-template-columns: 2fr 6fr 4fr 3fr 3fr 0.5fr;
        display: grid;
        grid-template-columns: var(--custom-grid-template-columns);
        grid-row: auto;

        justify-items: start;
        align-items: center;
        margin-bottom: auto;

        width: 900px;
        height: 35px;
        background: white;
        border-radius: 10px;
        /*border: 1px  rgb(212, 212, 212) solid;*/
        will-change: filter;

        box-shadow: 0 0 5px 0px rgb(200, 200, 200);

    }
    .book_row:hover{
        filter: drop-shadow(0 0 5px rgb(180, 180, 180));
    }
    .book_row_short{
        --custom-grid-template-columns: 2fr 6fr 4fr 3fr 3fr 0.5fr;
        display: grid;
        grid-template-columns: var(--custom-grid-template-columns);
        grid-row: auto;
        justify-items: start;
        align-items: center;
        margin-bottom: auto;

        width: 900px;
        height: 35px;
        background: rgb(255, 218, 218);
        border-radius: 10px;
        /*border: 1px  rgb(212, 212, 212) solid;*/
        will-change: filter;

        box-shadow: 0 0 30px 1px rgb(255, 0, 0);
    }
    .item-tag{
        align-self: center;
        justify-self: center;
        line-height: 0;
    }
    .item-title{
        align-self: center;
        justify-self: start;
        line-height: 0;
    }
    .item-author{
        align-self: center;
        justify-self: start;
        line-height: 0;
    }
    .item-edition{
        align-self: center;
        justify-self: start;
        line-height: 0;
    }
    .item-isbn{
        align-self: center;
        justify-self: end;
        line-height: 0;
        margin-right: 10px;
    }
</style>
