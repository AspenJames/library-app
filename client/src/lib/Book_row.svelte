<!--book row component-->
<script>
    import { createEventDispatcher } from 'svelte';
    import RowTag from './Row_tag.svelte';
    import DeleteButton from './Delete_button.svelte';
    export let title=undefined;
    export let author=undefined;
    export let edition=undefined;
    export let ISBN=undefined;
    export let read_status=false;

    const dispatch = createEventDispatcher();
  
    function forward(click) {
        dispatch('message', click.detail)};
</script>

<div class="book_row">
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
    <DeleteButton on:message={forward}></DeleteButton>
</div>

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

        box-shadow: -0.25px 0.5px 0px 1px rgb(220, 220, 220);

    }
    .book_row:hover{
        filter: drop-shadow(0 0 5px rgb(196, 196, 196));
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
