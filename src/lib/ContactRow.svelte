<script lang="ts">
    export let first: string;
    export let last: string;
    export let id: string;
    export let checked: boolean;
</script>

<tr>
    <th>
        <label>
            <input checked={checked} type="checkbox" class="checkbox" {id} />
        </label>
    </th>
    <td>
        <div class="font-bold">{first} {last}</div>
    </td>
</tr>
