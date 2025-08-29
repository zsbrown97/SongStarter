<script lang="ts">
    import GenerateButton from "$lib/components/GenerateButton.svelte";

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

<div class="flex column h-screen items-center justify-center">
    <GenerateButton
        onClick={getSongStart}
        text="Get Song Bones"
    />

    {#if songStart}
        <p><strong>Key Signature:</strong> {songStart.keySignature}</p><br/>
        <p><strong>Chords:</strong> {songStart.chords}</p><br/>
        <p><strong>Instrument:</strong> {songStart.instrument}</p>
    {/if}
</div>