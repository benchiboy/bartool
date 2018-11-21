// rcs_contract_mgr project doc.go

/*
rcs_contract_mgr document
*/
package main

/*
 <div id="+tblArray[p].table_name+">
	   <div class=\"note\" >
	        <div class=\"noteHeading\" style=\"display: flex;justify-content:flex-start\">
				<div style=\"display:inline-block;width:25%;\">
					<span>表名称:</span>
					<span id=\"table_name\">"+tblArray[p].table_name+"</span>
				</div>
				<div style=\"display:inline-block;width:25%;\">
					<span>描述:</span>
					<span id=\"table_desc\">"+tblArray[p].comment+"</span>
				</div>
				<div style=\"display:inline-block;width:25%;\">
					<span>数据条数(约):</span>
					<span id=\"table_rows\">"+tblArray[p].table_rows+"</span>
				</div>
				<div style=\"display:inline-block;width:25%;\">
					<span>是否加入迁移计划:</span>
					<span id=\"table_status\">"+tblArray[p].status+"</span>
				</div>
			</div>
	        <p class=\"notefooter\">
			 <ul class=\"menu\">
				<li role=\"presentation\">
	        	    <a role=\"menuitem\" tabindex=\"-1\" href=\"/complete/\">
	                <span class=\"glyphicon glyphicon-check\"></span></a>
	        	</li>
	        	<li role=\"presentation\">
	        	    <a role=\"menuitem\" tabindex=\"-1\" href=\"/trash/\">
	                <span class=\"glyphicon glyphicon-trash\"></span></a>
	        	</li>
	        	<li role=\"presentation\">
	        	    <a role=\"menuitem\" tabindex=\"-1\" href=\"/edit/\">
	                <span class=\"glyphicon glyphicon-pencil\"></span></a>
	        	</li>
	        	<li role=\"presentation\">
	              <a href=\"/category/\"> </a>
	        	  </li>
	        </ul>
			</p>
	    </div>
	</div >







SELECT * FROM information_schema.COLUMNS  WHERE TABLE_SCHEMA='crf_rcs_agent_db_ci_0' and COLUMN_KEY='PRI'

delete from 	agt_adjust_order	;
delete from 	agt_advance_flow	;
delete from 	agt_advance_plan	;
delete from 	agt_charge_flow	;
delete from 	agt_exempt_order	;
delete from 	agt_extend_loaninfo	;
delete from 	agt_extend_order	;
delete from 	agt_goods_flow	;
delete from 	agt_limit_flow	;
delete from 	agt_limit_flow_detail	;
delete from 	agt_limit_flow_detail_2017	;
delete from 	agt_loan_agreement	;
delete from 	agt_loan_agreement_2017	;
delete from 	agt_loan_agreement_detail	;
delete from 	agt_loan_agreement_detail_2017	;
delete from 	agt_loan_agreement_rule	;
delete from 	agt_loan_app	;
delete from 	agt_loan_app_201808	;
delete from 	agt_loan_goods	;
delete from 	agt_loan_info_ubike	;
delete from 	agt_loan_notarization	;
delete from 	agt_loan_order	;
delete from 	agt_loan_order_attachinfo	;
delete from 	agt_loan_order_ubike	;
delete from 	agt_loan_price	;
delete from 	agt_loan_relations	;
delete from 	agt_loan_times	;
delete from 	agt_mitigatesigned_info	;
delete from 	agt_offline_bankflow	;
delete from 	agt_offline_order	;
delete from 	agt_overflow_back_order	;
delete from 	agt_overflow_flow	;
delete from 	agt_pay_batch	;
delete from 	agt_pay_batch_detail	;
delete from 	agt_pay_flow	;
delete from 	agt_pay_flow_2017	;
delete from 	agt_repay_cancel_flow	;
delete from 	agt_repay_confirm_flow	;
delete from 	agt_repay_detail	;
delete from 	agt_repay_flow	;
delete from 	agt_repay_flow4notify	;
delete from 	agt_repay_flow_detail	;
delete from 	agt_repay_flow_goods	;
delete from 	agt_repay_offset_detail	;
delete from 	agt_repay_order	;
delete from 	agt_repay_order_2017	;
delete from 	agt_repay_order_detail	;
delete from 	agt_repay_order_detail_2017	;
delete from 	agt_repay_order_goods	;
delete from 	agt_repay_order_ubike	;
delete from 	agt_repay_plan	;
delete from 	agt_repay_plan_2017	;
delete from 	agt_repay_refund_flow	;
delete from 	agt_repay_remind_201703	;
delete from 	agt_repay_remind_201705	;
delete from 	agt_repay_remind_201706	;
delete from 	agt_repay_remind_201707	;
delete from 	agt_repay_remind_201708	;
delete from 	agt_repay_remind_201709	;
delete from 	agt_repay_remind_201710	;
delete from 	agt_repay_remind_201711	;
delete from 	agt_repay_remind_201712	;
delete from 	agt_repay_remind_201803	;
delete from 	agt_repay_seqcfg	;
delete from 	agt_repay_tryoffset	;
delete from 	agt_resp_log	;
delete from 	agt_return_goods	;
delete from 	agt_return_order	;
delete from 	agt_return_overflow_order	;
delete from 	agt_reward_batch	;
delete from 	agt_reward_detail	;
delete from 	agt_reward_detail_2017	;
delete from 	agt_sms_send	;
delete from 	agt_sms_sendlog	;
delete from 	agt_trxn_log	;
delete from 	agt_trxn_log_2017	;
delete from 	agt_trxn_log_2018	;
delete from 	agt_visit_log	;
delete from 	bas_agreement_group_cfg	;
delete from 	bas_channel_fee	;
delete from 	bas_error_code	;
delete from 	bas_merchant_info	;
delete from 	bas_merchant_product	;
delete from 	bas_mitigate_item	;
delete from 	bas_org_info	;
delete from 	bas_param_set	;
delete from 	bas_trxn_type	;
delete from 	txn_purchase_member_expand	;
delete from 	txn_purchase_order	;
delete from 	txn_purchase_order_detail	;
delete from 	txn_purchase_order_goods	;
delete from 	txn_purchase_refund	;
delete from 	txn_yearfee	;

		{Table_name: "agt_adjust_order", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_advance_flow", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_advance_plan", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_charge_flow", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_exempt_order", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_extend_loaninfo", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_extend_order", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_goods_flow", Key_id: "trxn_id", Update_mode: "I"},
		{Table_name: "agt_limit_flow", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_limit_flow_detail", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_loan_agreement", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_loan_agreement_detail", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_loan_agreement_rule", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_loan_app", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_loan_app_201808", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_loan_goods", Key_id: "id", Update_mode: "I"},
		{Table_name: "agt_loan_info_ubike", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_loan_notarization", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_loan_order", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_loan_order_attachinfo", Key_id: "id", Update_mode: "I"},
		{Table_name: "agt_loan_order_ubike", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_loan_price", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_loan_relations", Key_id: "id", Update_mode: "I"},
		{Table_name: "agt_loan_times", Key_id: "id", Update_mode: "I"},
		{Table_name: "agt_mitigatesigned_info", Key_id: "id", Update_mode: "I"},
		{Table_name: "agt_offline_bankflow", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_offline_order", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_overflow_back_order", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_overflow_flow", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_pay_batch", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_pay_batch_detail", Key_id: "auto_id", Update_mode: "I"},

		{Table_name: "agt_pay_flow", Key_id: "auto_id", Update_mode: "I",Migrate_Sync_Detail{Table_name:"agt_pay_flow",Key_no:"loan_no"}},



		{Table_name: "agt_repay_cancel_flow", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_repay_confirm_flow", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_repay_detail", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_repay_flow", Key_id: "trxn_id", Update_mode: "I"},
		{Table_name: "agt_repay_flow4notify", Key_id: "id", Update_mode: "I"},
		{Table_name: "agt_repay_flow_detail", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_repay_flow_goods", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_repay_offset_detail", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_repay_order", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_repay_order_detail", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_repay_order_goods", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_repay_order_ubike", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_repay_plan", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_repay_refund_flow", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_repay_remind_201703", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_repay_remind_201705", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_repay_remind_201706", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_repay_remind_201707", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_repay_remind_201708", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_repay_remind_201709", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_repay_remind_201710", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_repay_remind_201711", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_repay_remind_201712", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_repay_remind_201803", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_repay_seqcfg", Key_id: "id", Update_mode: "I"},
		{Table_name: "agt_repay_tryoffset", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_resp_log", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_return_goods", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_return_order", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_return_overflow_order", Key_id: "id", Update_mode: "I"},
		{Table_name: "agt_reward_batch", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_reward_detail", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_sms_send", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_sms_sendlog", Key_id: "auto_id", Update_mode: "I"},
		{Table_name: "agt_trxn_log", Key_id: "trxn_log_id", Update_mode: "I"},
		{Table_name: "agt_trxn_log_2017", Key_id: "trxn_log_id", Update_mode: "I"},
		{Table_name: "agt_visit_log", Key_id: "auto_id", Update_mode: "I"}}











*/
