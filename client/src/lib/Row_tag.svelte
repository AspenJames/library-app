<script>
	export let read_tag = false;
</script>

<button class:read_tag> <slot /> </button>

<style>
	button{
		/*default style is for unread*/
		--color: #ff8a00;
		--font-weight: bold;
		--border: 1px solid #ff8a00;  

		color: var(--color);
        font-weight: var(--font-weight);
		border: var(--border);
        background-color: white;
		width: 40px;
		font-size: 8px;
		height: 12px;
		border-radius: 2px;
		padding: 0px;
		text-align: center;
		font-weight: medium;
		display: inline-block;
	}
	button.read_tag{
		--color: grey;
		--font-weight: normal;
		--border: 1px solid grey;
	}
</style>