<script lang="ts">
	let { data } = $props();
</script>

<div class="min-h-screen bg-background flex items-center justify-center">
	<h1 class="text-6xl md:text-7xl font-bold text-text">Hello {data.username}</h1>
</div>
