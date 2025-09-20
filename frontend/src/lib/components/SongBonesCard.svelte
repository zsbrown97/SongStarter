<script lang="ts">
    import GenerateButton from "./GenerateButton.svelte";

    let songBones: {
        chords: string,
        instrument: string,
        keySignature: string,
        numerals: string,
    }

    async function getSongBones() {
        const res = await fetch('http://localhost:8080/api/songbones')
        songBones = await res.json()
    }
</script>

<div class="p-4 w-72 bg-white/10 rounded shadow">
    <div class="flex justify-center"> 
        <GenerateButton 
            onClick={getSongBones} text="Get Song Bones" 
        />
    </div>
    <div class="flex flex-col gap-4 mt-8">
        <div class="flex justify-between">
            <strong>Key Signature:</strong> 
            <span>{songBones?.keySignature}</span>
        </div>
        <div class="flex justify-between">
            <strong>Romans:</strong>
            <span>{songBones?.numerals}</span>
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