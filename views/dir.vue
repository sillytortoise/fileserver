<template>
	<el-cascader
		v-model="value"
		ref="mytree"
		:options="options"
		:props="{ expandTrigger: 'hover', checkStrictly: true }"
	></el-cascader>
</template>


<script>
module.exports = {
	data() {
		return {
			value: [],
			options: [],
		};
	},
	created: function () {
		axios.get(`/getdir`).then((res) => {
			this.options = [{ value: "/", label: "files", children: res.data }];
			console.log(this.options);
		});
	},
	methods: {
		handleNodeClick(data) {
			console.log(data);
		},
		getChecked() {
			return this.$refs.mytree.getCheckedNodes();
		},
	},
};
</script>