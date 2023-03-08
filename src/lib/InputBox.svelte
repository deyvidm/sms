<script lang="ts">
    import { afterUpdate } from "svelte";


    // this decorates the box if the input is wrong
    export function ValidateAndMark(): boolean{
        let val = Validate()
        Mark(val)
        return val
    }

    export function Mark(enabled: boolean) {
        if (enabled) {
            input.classList.remove("input-error")
        } else {
            input.classList.add("input-error")
        }
    }

    export function Validate(): boolean{
        return input.value.trim().length > 0
    }

    export let label: string;
    let input: HTMLInputElement;

    function keydown(){
        Mark(true)
    }

    export function Get(): string{
        return input.value
    }
</script>

<label for="" class="label">
    <span class="label-text">{label}</span>
</label>
<input
    bind:this={input}
    on:keydown={keydown}
    type="text"
    class="input input-bordered w-full max-w-xs"
/>