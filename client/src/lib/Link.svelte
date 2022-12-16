<script>
	export let link;
</script>

<a href={link} class={$$props.class}><slot></slot></a>

<style>
	a {
		color: var(--text-color);
        font-size: var(--font-size);
        font-weight: var(--font-weight);
	}
</style>