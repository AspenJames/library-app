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
    { read_status:false, title: 'This Indian Country', author: 'Frederick E. Hoxie', edition: 'undetermined', ISBN: '978-1-59420-365-7'},
    { read_status:undefined, title: undefined, author: undefined, edition: undefined, ISBN: undefined}
	];
</script>

<span class="lib_table_span">
  <div class="filter_search">
    <LibFilterMenu></LibFilterMenu>
    <div style="color:black">Search Bar</div>
  </div>
  <div class="button_colnames">
    <Button add_book_button button_text={"+"}></Button>
    <p style="color:black">Status, Title , Author , Edition , ISBN</p>
  </div>
  <div class="grid-container">
    {#each books as { read_status, title, author, edition, ISBN }, i}
      <BookRow read_status={read_status} title={title} author={author} edition={edition} ISBN={ISBN} ></BookRow>
    {/each}
  </div>
</span>

<style>
  .lib_table_span{
    display: inline-grid;
    grid-template-rows: 3;
    width: 900px;
  }
  .filter_search{
    display: inline-block;
  }
  .button_colnames{
    display: flex;
    flex-direction: row;
    justify-content: space-between;

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
</style>