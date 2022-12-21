<script>
	export let orange_button		= false;
	export let add_book_button		= false;
	export let wide_orange_button	= false;
	export let button_text 			= undefined;
</script>

<button class:orange_button class:wide_orange_button class:add_book_button on:click>{button_text}</button>

<style>
	button{
		/*default button style*/
		--background-color: white;
		--color: #ff8a00;
		--width: 120px;
		--height: 30px;
		--font-size: 14px;
		--font-weight: vold;
		--border: 2px solid #ff8a00;
		--border-radius: 8px;

		background-color: var(--background-color);
		color: var(--color);
		width: var(--width);
		font-size: var(--font-size);
		font-weight: var(--font-weight);
		border: var(--border);
		height: var(--height);
		border-radius: var(--border-radius);
		padding: 0px;
		text-align: center;
		font-weight: medium;
	}
	button.orange_button{
		--background-color: #ff8a00;
		--color: white;
		--font-weight: bold;
		will-change: filter;
	}
	button.orange_button:hover{
		filter: drop-shadow(0 0 5px #ff8a00);
	}
	button.wide_orange_button{
		--background-color: #ff8a00;
		--color: white;
		--width: 150px;
		--font-weight: bold;
		will-change: filter;
	}
	button.wide_orange_button:hover{
		filter: drop-shadow(0 0 5px #ff8a00)
	}
	button.add_book_button{
		--background-color: #ff8a00;
		--color: white;
		--width: 30px;
		--height: 30px;
		--font-size: 16px;
		--font-weight: bold;
		will-change: filter;
	}
	button.add_book_button:hover{
		filter: drop-shadow(0 0 5px #ff8a00)
	}
</style>

