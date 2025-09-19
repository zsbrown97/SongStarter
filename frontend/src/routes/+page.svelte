<script lang="ts">
    import GenerateButton from "$lib/components/GenerateButton.svelte";
    import SongBonesCard from "$lib/components/SongBonesCard.svelte";

    let songBones: {
        keySignature: string,
        chords: string,
        romans: string,
        instrument: string
    } | null = null

    async function getSongStart() {
        const res = await fetch('http://localhost:8080/api/songstart');
        songBones = await res.json()
    }
</script>

<div class="flex flex-col h-screen items-center justify-center">
    <SongBonesCard />
</div>