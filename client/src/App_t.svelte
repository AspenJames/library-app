<script>
  import Link from './lib/Link.svelte';
  import { slide, fade, fly } from 'svelte/transition';
  let visible=true;
</script>

<label>
	<input type="checkbox" bind:checked={visible}>
	visible
</label>

{#if visible}
  <div in:slide="{{ duration: 1000 }}" out:slide="{{ duration: 1000 }}">
    <Link link='#home' style='grid-column-start: 1;
    grid-column-end: four;
    grid-row-start: 2;
    grid-row-end: 3;'>
      Link
    </Link>
  </div>  
{/if}

<style >
  div{
    background-color: white;
    border-color: orange;
    width: 500px;
    height: 500px;
    color: black;
    font-size: 72px;
    display: grid;
    grid-template-columns: 1fr 2fr 1fr;
    grid-template-rows: repeat(4, 1fr);
  }
</style>