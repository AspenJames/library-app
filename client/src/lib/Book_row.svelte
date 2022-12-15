<!--book row component-->
<script>
    export let alt_row=false;
</script>

<div class="book_row" class:alt_row>
    <p>Hi</p>
</div>

<style>
    .book_row{
        width: 900px;
        height: 40px;
        background: white;
        border-radius: 10px;
        color: rgb(75, 75, 75);
        will-change: filter;
    }
    .book_row:hover{
        filter: drop-shadow(0 0 10px #005fdc);
    }
</style>