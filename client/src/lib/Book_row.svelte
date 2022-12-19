<!--book row component-->
<script>
    import { slide , fly} from 'svelte/transition';
    import RowTag from './Row_tag.svelte';
    import DeleteButton from './Delete_button.svelte';
    import ConfirmButton from './Confirm_button.svelte';
    export let title=undefined;
    export let author=undefined;
    export let edition=undefined;
    export let ISBN=undefined;
    export let read_status=false;


    let normalRowSize = true;
    function handleDeleteBookClick(delete_event) {
        normalRowSize=false;
        normalRowSize=normalRowSize;
    }
    function handleConfirmDeleteBookClick(confirm_delete_event) {
        normalRowSize=true;
        normalRowSize=normalRowSize;
    }
    
</script>

{#if normalRowSize}
    <div class="book_row" in:fly="{{ y:500, delay: 300, duration: 700 }}">
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
    <div class="book_row_short" out:fly="{{ x:-500, delay: 300, duration: 700 }}">
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
        <ConfirmButton on:message={handleConfirmDeleteBookClick}>CONFIRM</ConfirmButton>
    </div>
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

        box-shadow: 0 0 5px 0px rgb(220, 220, 220);

    }
    .book_row:hover{
        filter: drop-shadow(0 0 5px rgb(196, 196, 196));
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

        box-shadow: 0 0 20px 1px rgb(255, 0, 0);
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
