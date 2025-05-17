<script lang="ts">
	export let homeTeam: string;
	export let awayTeam: string;
	export let homeScore: number;
	export let awayScore: number;
	export let status: string;
</script>

<div class="mx-auto w-full max-w-md rounded-xl border bg-white p-4 shadow-md">
	<div class="mb-2 flex items-center justify-between text-sm text-gray-500">
		<span>{status}</span>
	</div>

	<div class="flex items-center justify-between text-lg font-semibold">
		<div>{awayTeam}</div>
		<div>{awayScore}</div>
	</div>

	<div class="mt-1 flex items-center justify-between text-lg font-semibold">
		<div>{homeTeam}</div>
		<div>{homeScore}</div>
	</div>
</div>
