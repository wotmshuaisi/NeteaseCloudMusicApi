/**
 * Created by ActingCute on 2017/8/18.
 */

$.fn.datetimepicker.dates['zh-CN'] = {
    days: ["星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六", "星期日"],
    daysShort: ["周日", "周一", "周二", "周三", "周四", "周五", "周六", "周日"],
    daysMin:  ["日", "一", "二", "三", "四", "五", "六", "日"],
    months: ["一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"],
    monthsShort: ["一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"],
    today: "今天",
    suffix: [],
    meridiem: ["上午", "下午"],
    rtl: true // 从右向左书写的语言你可以使用 rtl: true 来设置
};

$.extend($.validator.messages, {
    required: "必须填写",
    remote: "请修正此栏位",
    email: "请输入有效的电子邮件",
    url: "请输入有效的网址",
    date: "请输入有效的日期",
    dateISO: "请输入有效的日期 (YYYY-MM-DD)",
    number: "请输入正确的数字",
    digits: "只可输入数字",
    creditcard: "请输入有效的信用卡号码",
    equalTo: "你的输入不相同",
    extension: "请输入有效的后缀",
    maxlength: $.validator.format("最多 {0} 个字"),
    minlength: $.validator.format("最少 {0} 个字"),
    rangelength: $.validator.format("请输入长度为 {0} 至 {1} 之间的字串"),
    range: $.validator.format("请输入 {0} 至 {1} 之间的数值"),
    max: $.validator.format("请输入不大于 {0} 的数值"),
    min: $.validator.format("请输入不小于 {0} 的数值")
});

var msg = {
    invalid_username: "仅能使用字母、数字 和 中划线- 作为用户名，中划线 - 不能作为开头或结尾",
    invalid_ymd: "年月日有误,格式应该为xxxx-xx-xx"
}