export const fetchLiveGames = async () => {
	const res = await fetch('http://localhost:3000/nhl/games/live');
	if (!res.ok) throw new Error('Failed to fetch live games');
	return res.json();
};
