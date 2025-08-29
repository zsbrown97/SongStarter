<script lang="ts">
    let songStart: {
        keySignature: string,
        chords: string,
        instrument: string
    } | null = null

    async function getSongStart() {
        const res = await fetch('http://localhost:8080/api/songstart');
        songStart = await res.json()
        console.log(songStart)
    }
</script>

<button on:click={getSongStart}>Get Song Bones</button>

{#if songStart}
    <p><strong>Key Signature:</strong> {songStart.keySignature}</p>
    <p><strong>Chords:</strong> {songStart.chords}</p>
    <p><strong>Instrument:</strong> {songStart.instrument}</p>
{/if}