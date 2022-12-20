<script>
  import BookRow from './Book_row.svelte';
  import Button from './Button.svelte';
    import FilterAllButton from './filterAllButton.svelte';
    import FilterReadButton from './filterReadButton.svelte';
    import FilterUnreadButton from './filterUnreadButton.svelte';

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
    { read_status:false, title: 'This Indian Country', author: 'Frederick E. Hoxie', edition: 'undetermined', ISBN: '978-1-59420-365-7'},
    { read_status: true, title: 'Between The World And Me', author: 'Ta-Nehisi Coates', edition: 'First', ISBN: '978-0-8129-9354-7'},
    { read_status: false, title: 'The Beautiful Struggle', author: 'Ta-Nehisi Coates', edition: 'First', ISBN: '978-0-385-53746-0'},
    { read_status: false, title: 'Did Man Create God?', author: 'David E. Comings, M.D.', edition: 'First', ISBN: '978-1-878267-72-6'},
    { read_status: false, title: 'Gravity: how the weakest force in the universe shaped our lives', author: 'Brian Clegg', edition: 'First', ISBN: '978-0-312-61629-8'},
    { read_status: false, title: 'The China Study', author: 'T. Colin Campbell, PhD & Thomas M. Campbell II, MD', edition: 'First', ISBN: '978-1932100-66-2'},
    { read_status: false, title: "Quiet: the power of introverts in a world that can't stop talking", author: 'Susan Cain', edition: 'First', ISBN: '978-0-307-35215-6'},
    { read_status:true, title: 'Origins of Existence', author: 'Fred Adams', edition: 'First', ISBN: '978-1-5011-0008-6' },
		{ read_status:false, title: 'Clean Architecture', author: 'Robert C. Martin', edition: 'First', ISBN: '978-0-13-449416-6' },
		{ read_status:false, title: 'My Two Worlds', author: 'Sergio Chejfec', edition: 'First', ISBN: '978-1-934824-28-3' },
    { read_status:true, title: 'The Sixth Extinction', author: 'Elizabeth Kolbert', edition: 'First', ISBN: '978-1-250-06213-5'},
    { read_status:false, title: 'The Animals', author: 'Christian Kiefer', edition: 'NFS', ISBN: '978-0-87140-883-9'},
    { read_status:false, title: 'On The Road: The Original Scroll', author: 'Jack Kerouac', edition: 'deluxe', ISBN: '978-0-14-310546-6'},
    { read_status:false, title: 'Gratitude', author: 'Oliver Sacks', edition: 'First', ISBN: '978-0-451-49293-7'},
    { read_status:true, title: 'When Breath Becomes Air', author: 'Paul Kalanithi', edition: 'First', ISBN: '978-0-8129-8840-6'},
    { read_status:true, title: 'Amerika', author: 'Franz Kafka', edition: 'undetermined', ISBN: '978-0-8052-1064-4'},
    { read_status:false, title: 'This Indian Country', author: 'Frederick E. Hoxie', edition: 'undetermined', ISBN: '978-1-59420-365-7'},
    { read_status: true, title: 'Between The World And Me', author: 'Ta-Nehisi Coates', edition: 'First', ISBN: '978-0-8129-9354-7'},
    { read_status: false, title: 'The Beautiful Struggle', author: 'Ta-Nehisi Coates', edition: 'First', ISBN: '978-0-385-53746-0'},
    { read_status: false, title: 'Did Man Create God?', author: 'David E. Comings, M.D.', edition: 'First', ISBN: '978-1-878267-72-6'},
    
	];

  let inputTitle='title';
  let inputAuthor='author';
  let inputEdition='edition';
  let inputISBN='ISBN';

  let libraryTableVisible=true;
  let addBookCardVisible=false;
  let deleteBookCardVisible=false;

  function handleAddBookClick() {
    libraryTableVisible=false;
    addBookCardVisible=true;
    libraryTableVisible=libraryTableVisible;
    addBookCardVisible=addBookCardVisible;
  }
  function handleFinalAddBookClick() {
    addBookCardVisible=false;
    libraryTableVisible=true;
    addBookCardVisible=addBookCardVisible;
    libraryTableVisible=libraryTableVisible;
    //add function and if block for delete button bringing up the delete card
  }
  function handleDeleteBookClick() {
    deleteBookCardVisible=true;
    libraryTableVisible=false;
    deleteBookCardVisible=deleteBookCardVisible;
    libraryTableVisible=libraryTableVisible;
  };
  export let filterAll = true;
  export let filterUnread = false;
  export let filterRead = false;
  function handleFilterAllClick() {
    filterAll=true;
    filterUnread=false;
    filterRead=false;
    filterAll=filterAll;
    filterUnread=filterUnread;
    filterRead=filterRead;
  }
  function handleFilterUnreadClick() {
    filterAll=false;
    filterUnread=true;
    filterRead=false;
    filterAll=filterAll;
    filterUnread=filterUnread;
    filterRead=filterRead;
  }
  function handleFilterReadClick() {
    filterAll=false;
    filterUnread=false;
    filterRead=true;
    filterAll=filterAll;
    filterUnread=filterUnread;
    filterRead=filterRead;
  }

    //add function and if block for delete button bringing up the delete card

  function getFilteredBooks(books, var_unread, var_read){
    if(var_unread){
      return books.filter(book => {
        return book.read_status===false;
      });
    } else if(var_read) {
      return books.filter(book => {
        return book.read_status===true;
      });
    } else {
      return books
    }
};

</script>

<div class="lib_table_div">
  {#if libraryTableVisible}
    <h2 style="color:black; text-align:left;">My Library</h2>
    <div class='filter_menu'>
      <FilterAllButton filterAll={filterAll} on:filter-all={handleFilterAllClick}></FilterAllButton>
      <FilterUnreadButton filterUnread={filterUnread} on:filter-unread={handleFilterUnreadClick}></FilterUnreadButton>
      <FilterReadButton filterRead={filterRead} on:filter-read={handleFilterReadClick}></FilterReadButton>
    </div>
    <div class="grid-container">
      <div class="button_colnames_row">
        <div class="addBookButton">
          <Button add_book_button button_text={"+"} on:click={handleAddBookClick}></Button>
        </div>
        <p>Title</p>
        <p>Author</p>
        <p>Edition</p>
        <p class=isbn>ISBN</p>
      </div>
      <!--create another filtered list based on condition of LibFilterMenu-->
      {#if filterAll}
        {#each getFilteredBooks(books, false, false) as {read_status, title, author, edition, ISBN}}
          <BookRow read_status={read_status} title={title} author={author} edition={edition} ISBN={ISBN} on:message={handleDeleteBookClick}></BookRow>
        {/each}
      {:else if filterUnread}
        {#each getFilteredBooks(books, true, false) as {read_status, title, author, edition, ISBN}}
          <BookRow read_status={read_status} title={title} author={author} edition={edition} ISBN={ISBN} on:message={handleDeleteBookClick}></BookRow>
        {/each}
      {:else if filterRead}
        {#each getFilteredBooks(books, false, true) as {read_status, title, author, edition, ISBN}}
          <BookRow read_status={read_status} title={title} author={author} edition={edition} ISBN={ISBN} on:message={handleDeleteBookClick}></BookRow>
        {/each}
      {/if}
    </div>
  {:else if libraryTableVisible==false && addBookCardVisible}
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
      <div class="buttons" style="margin-top: 20px;"> !!!!!!!!on mouse click clear card and create aniamted checkmark verifying its added then return to library
          <Button button_text={"Add this book"} orange_button on:click={handleFinalAddBookClick}></Button>
      </div>
    </div>
  {/if}
</div>

<style>
  .lib_table_div{
    display: inline-grid;
    grid-template-rows: 3;
    width: 900px;
    height: 100%;
  }
  .filter_menu{
    height: 30px;
    width: 150px;
    display: flex;
    flex-direction: row;
    margin-bottom: 10px;
  }
  .button_colnames_row{
    grid-template-columns: 2fr 6fr 4fr 3fr 3fr 0.5fr;
    display: grid;
    grid-template-rows: auto;
    justify-items: start;
    align-items: bottom;
    color: grey;
    width: 100%;
  }
  .grid-container{
    --table-width: 900px;
    display: flex;
    flex-direction: column;
    align-items: center;
    row-gap: 10px;
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
    align-self: center;
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
  .isbn{
    justify-self: right;
    margin-right: 10px;
  }
</style>