<script>
	import { slide , fade } from 'svelte/transition';
	export let read_status = false;
	let startTransition=false;

	let buttonTransitionDelay=100;
	let buttonTransitionTime=100;

	function funcstartTransition() {
		startTransition=true;
		startTransition=startTransition;
	}

	function funcEndTransition() {
		startTransition=false;
		startTransition=startTransition;
	}
</script>
<!--first block for button state when in-transition from unread to read-->
{#if read_status && startTransition==false}
	<button class=read_button on:click={funcstartTransition} in:slide="{{ delay: buttonTransitionDelay, duration: buttonTransitionTime }}" out:slide="{{ delay: buttonTransitionDelay, duration: buttonTransitionTime }}" on:outroend="{() => read_status = false}" on:outroend={funcEndTransition}>READ</button>
{:else if read_status==false && startTransition==false}
	<button class=unread_button on:click={funcstartTransition}  in:slide="{{ delay: buttonTransitionDelay, duration: buttonTransitionTime }}" out:slide="{{ delay: buttonTransitionDelay, duration: buttonTransitionTime }}" on:outroend="{() => read_status = true}" on:outroend={funcEndTransition}>UNREAD</button>
	{/if}

<style>
	button{
        background-color: white;
		width: 45px;
		font-size: 9px;
		height: 15px;
		border-radius: 2px;
		padding: 0px;
		text-align: center;
		font-weight: medium;
	}
	button.unread_button{
		color: #ff8a00;
		font-weight: bold;
		border: 1px solid #ff8a00;

	}
	button.read_button{
		color: grey;
		font-weight: normal;
		border: 1px solid grey;
	}
</style>