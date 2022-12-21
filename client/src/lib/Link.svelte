<script>
	export let link;
	export let stylish=$$props.style;
</script>

<a href={link} class={$$props.class} style={stylish}><slot></slot></a>

<style>
	a {
		color: var(--text-color);
        font-size: var(--font-size);
        font-weight: var(--font-weight);
	}
</style>