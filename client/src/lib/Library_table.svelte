<script>
  import BookRow from './Book_row.svelte';
  import Button from './Button.svelte';
  import LibFilterMenu from './Lib_filter_menu.svelte';

  let books = [
		{ read_status:true, title: 'Origins of Existence', author: 'Fred Adams', edition: 'First', ISBN: '978-1-5011-0008-6' },
		{ read_status:false, title: 'Clean Architecture', author: 'Robert C. Martin', edition: 'First', ISBN: '978-0-13-449416-6' },
		{ read_status:false, title: 'My Two Worlds', author: 'Sergio Chejfec', edition: 'First', ISBN: '978-1-934824-28-3' },
    { read_status:true, title: 'The Sixth Extinction', author: 'Elizabeth Kolbert', edition: 'First', ISBN: '978-1-250-06213-5'},
    { read_status:false, title: 'The Animals', author: 'Christian Kiefer', edition: 'NFS', ISBN: '978-0-87140-883-9'},
    { read_status:false, title: 'On The Road: The Original Scroll', author: 'Jack Kerouac', edition: 'deluxe', ISBN: '978-0-14-310546-6'},
    { read_status:false, title: 'Gratitude', author: 'Oliver Sacks', edition: 'First', ISBN: '978-0-451-49293-7'},
    { read_status:true, title: 'When Breath Becomes Air', author: 'Paul Kalanithi', edition: 'First', ISBN: '978-0-8129-8840-6'},
    { read_status:true, title: 'Amerika', author: 'Franz Kafka', edition: 'undetermined', ISBN: '978-0-8052-1064-4'},
    { read_status:false, title: 'This Indian Country', author: 'Frederick E. Hoxie', edition: 'undetermined', ISBN: '978-1-59420-365-7'}
	];

  let inputTitle='title';
  let inputAuthor='author';
  let inputEdition='edition';
  let inputISBN='ISBN';

  export let libraryTableVisible=true;

  function handleAddBookClick() {
    libraryTableVisible=false;
    libraryTableVisible=libraryTableVisible;
  }
  function handleFinalAddBookClick() {
    libraryTableVisible=true;
    libraryTableVisible=libraryTableVisible;
  }

</script>

<div class="lib_table_span">
{#if libraryTableVisible}
  
    <h2 style="color:black; text-align:left;">My Library</h2>
    <div class='filter_menu'>
      <LibFilterMenu></LibFilterMenu>
    </div>
    <div class="button_colnames">
      <div class="addBookButton">
        <Button add_book_button button_text={"+"} on:click={handleAddBookClick}></Button>
      </div>
      <p>Title</p>
      <p>Author</p>
      <p>Edition</p>
      <p>ISBN</p>
    </div>
    <div class="grid-container">
      {#each books as { read_status, title, author, edition, ISBN }, i}
        <BookRow read_status={read_status} title={title} author={author} edition={edition} ISBN={ISBN} ></BookRow>
      {/each}
    </div>
  
{:else}
    <div class="card">
      <div class=cardP>
          <p>Add a book to your library</p>
      </div>
      <div class="inputs">
        <input bind:value={inputTitle}>
        <input bind:value={inputAuthor}>
        <input bind:value={inputEdition}>
        <input bind:value={inputISBN}>
      </div>
      <div class="buttons" style="margin-top: 20px;">
          <Button button_text={"Add this book"} orange_button on:click={handleFinalAddBookClick}></Button>
      </div>
    </div>
{/if}
</div>

<style>
  .lib_table_span{
    display: inline-grid;
    grid-template-rows: 3;
    width: 900px;
  }
  .filter_menu{
    height: 40px;
    width: 150px;
  }
  .button_colnames{
    display: grid;
    grid-template-columns: 1fr 4fr 3fr 1fr 3fr;
    grid-template-rows: auto;
    justify-items: start;
    align-items: bottom;
    color: grey;
  }
  .grid-container{
    --table-width: 900px;
    display: flex;
    flex-direction: column;
    align-items: center;
    row-gap: 10px;

    /*grid-template-columns: var(--table-width);
    grid-template-rows: repeat(10, 1fr);*/
    justify-content: space-around;

    background-color: white;
    font-family: "Montserrat", Sans-Serif;
    width: var(--table-width);
    height: auto;
    color: rgb(86, 86, 86);
    font-size: 15px;  
    font-weight: medium;
  }
  .card{
    justify-self: center;
    display: grid;
    grid-template-columns: auto;
    grid-template-rows: repeat(3, 1fr);
    width: 400px;
    height: 300px;
    border-radius: 30px;
    color: rgb(75, 75, 75);
    box-shadow: 0px 0px 15px 0px rgb(185, 185, 185);
  }
  .card p{
    line-height: 1.1;
    font-size: 40px;
    font-weight: bold;
    color: black;
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
</style>