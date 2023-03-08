<script lang="ts">
    import type pocketbaseEs from 'pocketbase';
    import { afterUpdate } from 'svelte';
    import { InputType } from './Types';

    export let label: string;
    let input: HTMLInputElement;

    export function Validate(): boolean {
        let dig = getDigits(input.value);
        if (dig.length != 10) {
            return false;
        }
        return true;
    }

    export function ValidateAndMark(): boolean {
        let val = Validate();
        Mark(val);
        return val;
    }

    export function Mark(enabled: boolean) {
        if (enabled) {
            input.classList.remove('input-error');
        } else {
            input.classList.add('input-error');
        }
    }

    export function Get(): string {
        return "+1"+getDigits(input.value).join("");
    }

    function getDigits(val: string): Number[] {
        // pluck out every digit from the input string
        let ret: Array<Number> = [];
        const regex = /\d/g;
        val.match(regex)?.forEach((val, i) => {
            ret.push(Number(val));
        });
        return ret;
    }

    // this function formats the value inside the text box to look like:
    // (123)-456-789
    function formatUserinput(rawVal: string): string {
        if (isNaN(parseInt(rawVal[rawVal.length-1]))) {
            rawVal = rawVal.slice(0,rawVal.length-2);
        }
        let digits = getDigits(rawVal);
        let displayString: string[] = [];
        for (let i = 0; i < digits.length; i++) {
            if (i == 0) {
                displayString.push('(');
            }
            if (i > 9) {
                break;
            }
            displayString.push(digits[i].toString());
            if (i == 2) {
                displayString.push(')');
            }
            if (i == 2 || i == 5) {
                displayString.push('-');
            }
        }
        return displayString.join('');
    }


    function onkeydown() {
        Mark(true);
        let formattedInput = formatUserinput(input.value);
        input.value = formattedInput;
    }
</script>

<label for="" class="label">
    <span class="label-text">{label}</span>
</label>
<input
    bind:this={input}
    on:input={onkeydown}
    type="text"
    class="input input-bordered w-full max-w-xs"
/>
