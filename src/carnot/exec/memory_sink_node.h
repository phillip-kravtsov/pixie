#pragma once

#include <memory>
#include <string>
#include <vector>

#include "src/carnot/exec/exec_node.h"
#include "src/carnot/exec/exec_state.h"

namespace pl {
namespace carnot {
namespace exec {

class MemorySinkNode : public SinkNode {
 public:
  MemorySinkNode() = default;
  std::string TableName() const { return plan_node_->TableName(); }

 protected:
  std::string DebugStringImpl() override;
  Status InitImpl(const plan::Operator &plan_node, const schema::RowDescriptor &output_descriptor,
                  const std::vector<schema::RowDescriptor> &input_descriptors) override;
  Status PrepareImpl(ExecState *exec_state) override;
  Status OpenImpl(ExecState *exec_state) override;
  Status CloseImpl(ExecState *exec_state) override;
  Status ConsumeNextImpl(ExecState *exec_state, const schema::RowBatch &rb) override;

 private:
  std::unique_ptr<plan::MemorySinkOperator> plan_node_;
  std::unique_ptr<schema::RowDescriptor> input_descriptor_;
  std::shared_ptr<schema::Table> table_;
};

}  // namespace exec
}  // namespace carnot
}  // namespace pl
