<script lang="ts">
	import { onMount } from 'svelte';
	import GameCard from '$components/GameCard.svelte';
	import { fetchLiveGames } from '$lib/api';

	type Game = {
		gameID: number;
		homeTeam: string;
		awayTeam: string;
		homeScore: number;
		awayScore: number;
		status: string;
	};

	let games: Game[] = [];
	let error: string | null = null;

	const fetchGames = async () => {
		try {
			const data = await fetchLiveGames();
			// if (!res.ok) throw new Error(`HTTP ${res.status}`);
			// const data = await res.json();
			games = data.games;
		} catch (e) {
			console.error('Fetch error:', e);
			error = 'Failed to load games';
		}
	};

	// Auto-refresh every 30 seconds
	onMount(() => {
		fetchGames();
		const interval = setInterval(fetchGames, 30000);
		return () => clearInterval(interval);
	});
</script>

{#if error}
	<p class="text-red-500">{error}</p>
{:else if games.length === 0}
	<p class="text-gray-500">No games available.</p>
{:else}
	<div class="mt-6 space-y-4">
		{#each games as game}
			<GameCard
				homeTeam={game.homeTeam}
				awayTeam={game.awayTeam}
				homeScore={game.homeScore}
				awayScore={game.awayScore}
				status={game.status}
			/>
		{/each}
	</div>
{/if}
