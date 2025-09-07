import React, { useRef, useEffect } from "react";
import * as echarts from "echarts";

const MyChart: React.FC = () => {
	const chartRef = useRef<HTMLDivElement | null>(null); // 类型注解

	useEffect(() => {
		if (chartRef.current) {
			const chart = echarts.init(chartRef.current); // 初始化图表

			const options: echarts.EChartOption = {
				// 使用 EChartOption 类型
				title: {
					text: "ECharts Example",
				},
				tooltip: {},
				xAxis: {
					data: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
				},
				yAxis: {},
				series: [
					{
						name: "Sales",
						type: "line",
						data: [5, 20, 36, 10, 10, 20, 30],
					},
				],
			};

			// 使用配置项和数据显示图表
			chart.setOption(options);

			// 在组件卸载时销毁图表实例，防止内存泄漏
			return () => {
				chart.dispose();
			};
		}
	}, []);

	return <div ref={chartRef} style={{ width: "100%", height: "400px" }} />;
};

export default MyChart;
