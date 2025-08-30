<script lang="ts">
    import GenerateButton from "$lib/components/GenerateButton.svelte";

    let songBones: {
        keySignature: string,
        chords: string,
        instrument: string
    } | null = null

    async function getSongStart() {
        const res = await fetch('http://localhost:8080/api/songstart');
        songBones = await res.json()
    }
</script>

<div class="flex flex-col h-screen items-center justify-center">
    <div class="p-4 w-72 bg-white/10 rounded shadow">
        <GenerateButton onClick={getSongStart} text="Get Song Bones" />

        <div class="flex flex-col gap-4 mt-8">
            <div class="flex justify-between">
                <strong>Key Signature:</strong> 
                <span>{songBones?.keySignature}</span>
            </div>
            <div class="flex justify-between">
                <strong>Chords:</strong> 
                <span>{songBones?.chords}</span>
            </div>
            <div class="flex justify-between">
                <strong>Instrument:</strong>
                <span>{songBones?.instrument}</span>
            </div>
        </div>
    </div>
</div>