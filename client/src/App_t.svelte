<script>
  import Link from './lib/Link.svelte';
</script>

<div class='parent'>
  <Link link={'https://www.google.com'} class='child'>
    it's me
  </Link>
</div>

<style>
  .parent :global(.child){
    color: lime;
    font-size: 72px;
  }
</style>