<?php
$key="AAAvw3YlqoEk6G4HqRKDAYpHKZWxBBB";
$iv="AAAC1FplieBBB"; 
$mid="MS12345678";

// $data1="MerchantID=MS12345678&TimeStamp=1663040304&Version=2.0&RespondType=String&MerchantOrderNo=Vanespl_ec_1663040304&Amt=30&NotifyURL=https%3A%2F%2Fwebhook.site%2Fd4db5ad1-2278-466a-9d6678585c0dbadb&ReturnURL=&ItemDesc=test";
$data1=http_build_query(array( 'MerchantID'=>$mid, 'TimeStamp'=>time(), 'Version'=>'2.0', 'RespondType'=>'String', 'MerchantOrderNo'=>'Vanespl_ec_'.time(), 'Amt'=>'30', 'NotifyURL'=>'https://webhook.site/d4db5ad1-2278-466a-9d6678585c0dbadb', 'ReturnURL'=>'', 'ItemDesc'=>'test', ));
// $edata1=bin2hex(openssl_encrypt($data1, "AES-256-CBC", $key, OPENSSL_RAW_DATA, $iv));

// echo $edata1;

// $data1=http_build_query(
//   array( 
//     'MerchantID'=>$mid, 
//     'TimeStamp'=>'1663040304', 
//     'Version'=>'2.0', 
//     'RespondType'=>'String', 
//     'MerchantOrderNo'=>'Vanespl_ec_'.'1663040304', 
//     'Amt'=>'30', 
//     'NotifyURL'=>'https://webhook.site/d4db5ad1-2278-466a-9d6678585c0dbadb', 
//     'ReturnURL'=>'', 
//     'ItemDesc'=>'test', 
//   ));

$edata1=bin2hex(openssl_encrypt($data1, "AES-256-CBC", $key, OPENSSL_RAW_DATA, $iv));

echo $edata1;
?>
