<script lang="ts">
    export let first: string;
    export let last: string;
    export let id: string;
    export let checked: boolean;
    import { beforeUpdate, createEventDispatcher, onMount } from 'svelte';

    const dispatch = createEventDispatcher();

    function onClick() {
        checked = !checked;
        dispatch('message', {
            checked: checked,
            id: id,
        });
    }
</script>

<tr on:click={onClick}>
    <td>
        <label>
            <input bind:checked {id} type="checkbox" class="checkbox" />
        </label>
    </td>
    <td>
        <div class="font-bold">{first} {last}</div>
    </td>
</tr>
